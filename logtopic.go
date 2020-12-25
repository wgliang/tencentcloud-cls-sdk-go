package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LogTopics struct {
	LogSets []LogSet `json:"topics"`
}

type LogTopic struct {
	LogTopicID     string `json:"topic_id"`
	LogTopicName   string `json:"topic_name"`
	LogSetID       string `json:"logset_id"`
	SQLFlag        bool   `json:"sql_flag"`
	Shipper        bool   `json:"shipper"`
	Path           string `json:"path"`
	PartitionCount int    `json:"partition_count"`
	// MultiWildPath  []byte `json:"multi_wild_path"` //TODO
	// MaxSplit int    `json:"maxSplit"` //TODO
	LogType string `json:"log_type"`
	// LogFormat string `json:"log_format"` //TODO
	Isolated int  `json:"isolated"`
	Index    bool `json:"index"`
	// ExtractRule   ExtractRule `json:"extract_rule"` //TODO
	CreateTime    string `json:"create_time"`
	Collection    bool   `json:"collection"`
	AutoPartition bool   `json:"autoPartition"`
	// ExcludePaths   []byte      `json:"ExcludePaths"` //TODO
}

type ExtractRule struct {
	AutoPartition []byte `json:"filter_keys"`
	ExcludePaths  []byte `json:"filter_regex"`
}

// TODO
func (cls *ClSCleint) CreateLogTopic(logTopic LogTopic) (string, error) {
	data, err := json.Marshal(logTopic)
	if err != nil {
		return "", err
	}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer(data)
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/topic", nil, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/topic", cls.Host), body)
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
	topic := LogTopic{}
	bod, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bod, &topic); err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(topic)
	return topic.LogTopicID, nil
}

func (cls *ClSCleint) GetLogTopic(logTopicID string) (logTopic LogTopic, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", logTopicID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/topic", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/topic?topic_id=%s", cls.Host, logTopicID), nil)
	if err != nil {
		return logTopic, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return logTopic, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &logTopic); err != nil {
		fmt.Println(err)
		return logTopic, err
	}

	return logTopic, nil
}

func (cls *ClSCleint) GetLogTopics(logsetId string) (logTopics LogTopics, err error) {
	var params = url.Values{"logset_id": {fmt.Sprintf("%s", logsetId)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/topics", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/topics?logset_id=%s", cls.Host, logsetId), nil)
	if err != nil {
		return logTopics, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return logTopics, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &logTopics); err != nil {
		return logTopics, err
	}

	return logTopics, nil
}

func (cls *ClSCleint) UpdateLogTopic(logTopic LogTopic) error {
	data, err := json.Marshal(logTopic)
	if err != nil {
		return err
	}
	// data := fmt.Sprintf("{\"topic_id\":\"%s\",\"topic_name\":\"%s\"}", logSetID, logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer(data)
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/topic", nil, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/topic", cls.Host), body)
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

func (cls *ClSCleint) DeleteLogTopic(logTopicID string) error {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", logTopicID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"DELETE", "/topic", params, headers, 300)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/topic?topic_id=%s", cls.Host, logTopicID), nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}
