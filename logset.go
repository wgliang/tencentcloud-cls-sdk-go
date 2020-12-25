package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type LogSets struct {
	LogSets []LogSet `json:"logsets"`
}

type LogSet struct {
	LogSetID     string `json:"logset_id"`
	LogSetName   string `json:"logset_name"`
	CreateTime   string `json:"create_time"`
	Period       int    `json:"period"`
	TopicsNumber int    `json:"topics_number"`
}

func (cls *ClSCleint) CreateLogSet(logSetName string, period int) (string, error) {
	data := fmt.Sprintf("{\"logset_name\":\"%s\",\"period\":%d}", logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/logset", nil, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/logset", cls.Host), body)
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
	logSet := LogSet{}
	bod, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bod, &logSet); err != nil {
		fmt.Println(err)
		return "", err
	}

	return logSet.LogSetID, nil
}

func (cls *ClSCleint) GetLogSet(logSetID string) (logSet LogSet, err error) {
	var params = url.Values{"logset_id": {fmt.Sprintf("%s", logSetID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/logset", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/logset?logset_id=%s", cls.Host, logSetID), nil)
	if err != nil {
		return logSet, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return logSet, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &logSet); err != nil {
		fmt.Println(err)
		return logSet, err
	}

	return logSet, nil
}

func (cls *ClSCleint) GetLogSets() (logSets LogSets, err error) {
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/logsets", nil, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/logsets", cls.Host), nil)
	if err != nil {
		return logSets, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return logSets, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &logSets); err != nil {
		return logSets, err
	}

	return logSets, nil
}

func (cls *ClSCleint) UpdateLogSet(logSetID, logSetName string, period int) error {
	data := fmt.Sprintf("{\"logset_id\":\"%s\",\"logset_name\":\"%s\",\"period\":%d}", logSetID, logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/logset", nil, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/logset", cls.Host), body)
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

func (cls *ClSCleint) DeleteLogSet(logSetID string) error {
	var params = url.Values{"logset_id": {fmt.Sprintf("%s", logSetID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"DELETE", "/logset", params, headers, 300)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/logset?logset_id=%s", cls.Host, logSetID), nil)
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(string(body))

	return nil
}
