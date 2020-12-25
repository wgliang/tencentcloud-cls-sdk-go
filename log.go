package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/proto"
)

type Cursor struct {
	Value string `json:"cursor"`
}

func (cls *ClSCleint) UploadLog(logTopicID string, logGroupList LogGroupList, hash string, compress bool) error {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", logTopicID)}}

	data, err := proto.Marshal(&logGroupList)
	if err != nil {
		return err
	}

	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/structuredlog", params, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/structuredlog?topic_id=%s", cls.Host, logTopicID), body)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))
	req.Header.Add("Content-Type", "application/x-protobuf")
	if hash != "" {
		req.Header.Add("x-cls-hashkeye", hash)
	}
	if compress {
		req.Header.Add("x-cls-compress-type", "lz4")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("%d", resp.StatusCode)
	}

	return nil
}

func (cls *ClSCleint) GetLogStart(logTopicID, start string) (cursor string, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", logTopicID)}, "start": {fmt.Sprintf("%s", start)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/cursor", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/cursor?topic_id=%s&start=%s", cls.Host, logTopicID, url.QueryEscape(start)), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%d", resp.StatusCode)
	}

	cursorStruct := Cursor{}
	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &cursorStruct); err != nil {
		fmt.Println(err)
		return "", err
	}

	return cursorStruct.Value, nil
}

// TODO 一直返回400
func (cls *ClSCleint) SearchLog(requestDataMap map[string]string) (string, error) {
	var params = url.Values{}
	var urlString string
	for k, v := range requestDataMap {
		params.Add(fmt.Sprintf("%s", k), fmt.Sprintf("%s", v))
		urlString = fmt.Sprintf("%s&%s=%s", urlString, k, url.QueryEscape(v))
	}
	urlString = urlString[1:len(urlString)]
	fmt.Println(fmt.Sprintf("http://%s/searchlog?%s", cls.Host, urlString))
	// var params = url.Values{"logset_id": {fmt.Sprintf("%s", logSetID)},
	// 	"topic_ids": {fmt.Sprintf("%s", topicIDs)}, "start_time": {fmt.Sprintf("%s", startTime)},
	// 	"end_time": {fmt.Sprintf("%s", endTime)}, "query_string": {fmt.Sprintf("%s", queryString)},
	// 	"limit": {fmt.Sprintf("%s", limit)}, "context": {fmt.Sprintf("%s", context)},
	// 	"sort": {fmt.Sprintf("%s", sort)},
	// }
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/searchlog", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/searchlog?%s", cls.Host, urlString), nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println(string(body))

	return "", nil
}

func (cls *ClSCleint) DowloadLog(logTopicID, cursor, count string) error {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", logTopicID)}, "cursor": {fmt.Sprintf("%s", cursor)}, "count": {fmt.Sprintf("%s", count)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/log", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/log?topic_id=%s&cursor=%s&count=%s", cls.Host, logTopicID, url.QueryEscape(cursor), count), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("%d", resp.StatusCode)
	}

	return nil
}
