package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ShipperList struct {
	Shippers []Shipper `json:"shippers"`
}
type Shipper struct {
	TopicID     string   `json:"topic_id"`
	Bucket      string   `json:"bucket"`
	Prefix      string   `json:"prefix"`
	ShipperName string   `json:"shipper_name"`
	Interval    int      `json:"interval"`
	MaxSize     int      `json:"max_size"`
	Partition   string   `json:"partition"`
	Compress    Compress `json:"compress"`
	Content     Content  `json:"content"`
}

type ShipperID struct {
	ShipperID string `json:"shipper_id"`
}

type Compress struct {
	Format string `json:"format"`
}

type Content struct {
	Format  string  `json:"format"`
	CsvInfo CsvInfo `json:"csv_info"`
}

type CsvInfo struct {
	PrintKey         string `json:"print_key"`
	Keys             string `json:"keys"`
	Delimiter        string `json:"delimiter"`
	EscapeChar       string `json:"escape_char"`
	NonExistingField string `json:"non_existing_field"`
}

type Tasks struct {
	Tasks []TaskInfo `json:"tasks"`
}

type TaskInfo struct {
	TaskID     string `json:"task_id"`
	ShipperID  string `json:"shipper_id"`
	TopicID    string `json:"topic_id"`
	RangeStart string `json:"range_start"`
	RangeEnd   string `json:"range_end"`
	StartTime  string `json:"start_time"`
	EndTime    string `json:"end_time"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

// CreateShipper 创建新的投递任务，如果使用此接口，需要自行处理 CLS 对指定 Bucket 的写权限。
func (cls *ClSCleint) CreateShipper(shipper *Shipper) (string, error) {
	data, err := json.Marshal(shipper)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	// data := fmt.Sprintf("{\"logset_name\":\"%s\",\"period\":%d}", logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/shipper", nil, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/shipper", cls.Host), body)
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

	shipperID := ShipperID{}
	bod, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bod, &shipperID); err != nil {
		fmt.Println(err)
		return "", err
	}

	return shipperID.ShipperID, nil
}

// GetShipper 本接口用于获取指定投递策略的详细信息
func (cls *ClSCleint) GetShipper(shipperID string) (shipper Shipper, err error) {
	var params = url.Values{"shipper_id": {fmt.Sprintf("%s", shipperID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/shipper", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/shipper?shipper_id=%s", cls.Host, shipperID), nil)
	if err != nil {
		return shipper, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return shipper, err
	}
	if resp.StatusCode != 200 {
		return shipper, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &shipper); err != nil {
		fmt.Println(err)
		return shipper, err
	}

	return shipper, nil
}

// GetTopicShipperList 本接口用于获取指定日志主题的投递策略详细列表
func (cls *ClSCleint) GetTopicShipperList(topicID string) (shipperList ShipperList, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/shippers", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/shippers?topic_id=%s", cls.Host, topicID), nil)
	if err != nil {
		return shipperList, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return shipperList, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &shipperList); err != nil {
		fmt.Println(err)
		return shipperList, err
	}

	return shipperList, nil
}

// GetShipperList 本接口可用于获取投递任务信息列表。
func (cls *ClSCleint) GetShipperList(shipperID, startTime, endTime string) (tasks Tasks, err error) {
	var params = url.Values{"shipper_id": {fmt.Sprintf("%s", shipperID)}, "start_time": {fmt.Sprintf("%s", startTime)}, "end_time": {fmt.Sprintf("%s", endTime)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/tasks", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/tasks?shipper_id=%s&start_time=%s&end_time=%s", cls.Host, shipperID, url.QueryEscape(startTime), url.QueryEscape(endTime)), nil)
	if err != nil {
		return tasks, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return tasks, err
	}

	if resp.StatusCode != 200 {
		return tasks, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &tasks); err != nil {
		fmt.Println(err)
		return tasks, err
	}

	return tasks, nil
}

// UpdateShipper 本接口可用于修改现有的投递任务，客户如果使用此接口，需要自行处理 CLS 对指定 Bucket 的写权限。
func (cls *ClSCleint) UpdateShipper(shipper *Shipper) (string, error) {
	data, err := json.Marshal(shipper)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	// data := fmt.Sprintf("{\"logset_name\":\"%s\",\"period\":%d}", logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/shipper", nil, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/shipper", cls.Host), body)
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

	shipperID := ShipperID{}
	bod, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bod, &shipperID); err != nil {
		fmt.Println(err)
		return "", err
	}

	return shipperID.ShipperID, nil
}

// UpdateTask 本接口可用于重试失败的投递任务。
func (cls *ClSCleint) UpdateTask(shipper *Shipper) error {
	data, err := json.Marshal(shipper)
	if err != nil {
		return err
	}
	fmt.Println(string(data))
	// data := fmt.Sprintf("{\"logset_name\":\"%s\",\"period\":%d}", logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/task", nil, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/task", cls.Host), body)
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

// DeleteShipper 本接口用于删除投递配置。
func (cls *ClSCleint) DeleteShipper(shipperID string) error {
	var params = url.Values{"shipper_id": {fmt.Sprintf("%s", shipperID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"DELETE", "/shipper", params, headers, 300)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/shipper?shipper_id=%s", cls.Host, shipperID), nil)
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
