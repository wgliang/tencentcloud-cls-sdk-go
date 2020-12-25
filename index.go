package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Index struct {
	TopicID   string `json:"topic_id"`
	Effective bool   `json:"effective"`
	Rule      Rule   `json:"rule"`
}

type Rule struct {
	FullText FullText `json:"full_text"`
	KeyValue KeyValue `json:"key_value"`
}
type FullText struct {
	CaseSensitive bool `json:"case_sensitive"`
}

type KeyValue struct {
	CaseSensitive bool     `json:"case_sensitive"`
	Keys          []string `json:"keys"`
	Types         []string `json:"types"`
}

func (cls *ClSCleint) GetLogIndex(logTopicID string) (index Index, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", logTopicID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/index", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/index?topic_id=%s", cls.Host, logTopicID), nil)
	if err != nil {
		return index, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return index, err
	}

	if resp.StatusCode != 200 {
		return index, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &index); err != nil {
		return index, err
	}

	return index, nil
}

func (cls *ClSCleint) UpdateLogIndex(logIndex Index) error {
	data, err := json.Marshal(logIndex)
	if err != nil {
		return err
	}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer(data)
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/index", nil, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/index", cls.Host), body)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))
	req.Header.Add("Content-Type", "application/json")

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
