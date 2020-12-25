// Writing a basic HTTP server is easy using the
// `net/http` package.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	cls "github.com/wgliang/tencentcloud-cls-sdk-go"
)

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func index(w http.ResponseWriter, req *http.Request) {
	indexMaps := map[string]cls.Index{
		"topic01": {
			TopicID:   "topic01",
			Effective: false,
			Rule: cls.Rule{
				FullText: cls.FullText{
					CaseSensitive: false,
				},
				KeyValue: cls.KeyValue{
					CaseSensitive: false,
					Keys:          []string{"key1", "key2"},
					Types:         []string{"type1", "type2"},
				},
			},
		},
		"topic03": {
			TopicID:   "topic03",
			Effective: false,
			Rule: cls.Rule{
				FullText: cls.FullText{
					CaseSensitive: false,
				},
				KeyValue: cls.KeyValue{
					CaseSensitive: false,
					Keys:          []string{"key1", "key2"},
					Types:         []string{"type1", "type2"},
				},
			},
		},
	}

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	switch req.Method {
	case "GET":
		topicID := req.FormValue("topic_id")
		if v, ok := indexMaps[topicID]; ok {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
				http.Error(w, "data is error", http.StatusBadRequest)
			}
			fmt.Fprintf(w, string(b))
			return
		} else {
			http.Error(w, "Not find topic.", http.StatusNotFound)
			return
		}
	case "PUT":
		var index cls.Index
		err := json.NewDecoder(req.Body).Decode(&index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, ok := indexMaps[index.TopicID]; ok {
			fmt.Fprintf(w, "200")
			return
		} else {
			http.Error(w, "Not find topic.", http.StatusNotFound)
			return
		}
	default:
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}

func structuredlog(w http.ResponseWriter, req *http.Request) {
	indexMaps := map[string]cls.Index{
		"topic01": {
			TopicID:   "topic01",
			Effective: false,
			Rule: cls.Rule{
				FullText: cls.FullText{
					CaseSensitive: false,
				},
				KeyValue: cls.KeyValue{
					CaseSensitive: false,
					Keys:          []string{"key1", "key2"},
					Types:         []string{"type1", "type2"},
				},
			},
		},
		"topic03": {
			TopicID:   "topic03",
			Effective: false,
			Rule: cls.Rule{
				FullText: cls.FullText{
					CaseSensitive: false,
				},
				KeyValue: cls.KeyValue{
					CaseSensitive: false,
					Keys:          []string{"key1", "key2"},
					Types:         []string{"type1", "type2"},
				},
			},
		},
	}

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	switch req.Method {
	case "GET":
		topicID := req.FormValue("topic_id")
		if v, ok := indexMaps[topicID]; ok {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
				http.Error(w, "data is error", http.StatusBadRequest)
			}
			fmt.Fprintf(w, string(b))
			return
		} else {
			http.Error(w, "Not find topic.", http.StatusNotFound)
			return
		}
	case "PUT":
		var index cls.Index
		err := json.NewDecoder(req.Body).Decode(&index)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if _, ok := indexMaps[index.TopicID]; ok {
			fmt.Fprintf(w, "200")
			return
		} else {
			http.Error(w, "Not find topic.", http.StatusNotFound)
			return
		}
	case "POST":
		topicID := req.FormValue("topic_id")
		if v, ok := indexMaps[topicID]; ok {
			b, err := json.Marshal(v)
			if err != nil {
				fmt.Println("error:", err)
				http.Error(w, "data is error", http.StatusBadRequest)
			}
			fmt.Fprintf(w, string(b))
			return
		} else {
			http.Error(w, "Not find topic.", http.StatusNotFound)
			return
		}
	default:
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
}

func headers(w http.ResponseWriter, req *http.Request) {
	// This handler does something a little more
	// sophisticated by reading all the HTTP request
	// headers and echoing them into the response body.
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// We register our handlers on server routes using the
	// `http.HandleFunc` convenience function. It sets up
	// the *default router* in the `net/http` package and
	// takes a function as an argument.
	http.HandleFunc("/index", index)
	// http.HandleFunc("/log", log)
	http.HandleFunc("/structuredlog", structuredlog)

	// http.HandleFunc("/logset", logset)
	// http.HandleFunc("/topic", topic)
	// http.HandleFunc("/log", log)

	// Finally, we call the `ListenAndServe` with the port
	// and a handler. `nil` tells it to use the default
	// router we've just set up.
	http.ListenAndServe(":8080", nil)
}
