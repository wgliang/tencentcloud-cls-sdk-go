package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ConsumerGroups struct {
	Machines []Machine `json:"consumer_groups"`
}

type ConsumerGroup struct {
	ConsumerGroup string `json:"consumer_group"`
	Order         bool   `json:"order"`
	Timeout       int    `json:"timeout"`
}

type ConsumerCursor struct {
	Cursor string `json:"cursor"`
}

type ConsumerGroupCursors struct {
	Cursors string `json:"cursors"`
}

type ConsumerGroupCursor struct {
	ConsumerID  string `json:"consumer_id"`
	Cursor      string `json:"cursor"`
	PartitionID string `json:"partition_id"`
	UpdateTime  string `json:"update_time"`
}

type ConsumerHeartbeat struct {
	ConsumerGroup   string   `json:"consumer_group"`
	ConsumerID      string   `json:"consumer_id"`
	PartitionIDList []string `json:"partition_id_list"`
}

// TODO
func (cls *ClSCleint) CreateConsumerGroup(group *ConsumerGroup, topicID string) error {
	data, err := json.Marshal(group)
	if err != nil {
		return err
	}

	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}}

	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/consumergroup", params, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/consumergroup", cls.Host), body)
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

func (cls *ClSCleint) GetConsumerCursor(topicID, partitionID, from string) (cursor ConsumerCursor, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}, "partition_id": {fmt.Sprintf("%s", partitionID)}, "from": {fmt.Sprintf("%s", from)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/machinegroup", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/machinegroup?topic_id=%s&partition_id=%s&from=%s", cls.Host, topicID, partitionID, from), nil)
	if err != nil {
		return cursor, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cursor, err
	}
	if resp.StatusCode != 200 {
		return cursor, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &cursor); err != nil {
		return cursor, err
	}

	return cursor, nil
}

func (cls *ClSCleint) GetConsumerGroupCursors(topicID, partitionID, consumerGroup string) (cursors ConsumerGroupCursors, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}, "partition_id": {fmt.Sprintf("%s", partitionID)}, "consumer_group": {fmt.Sprintf("%s", consumerGroup)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/consumergroupcursor", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/consumergroupcursor?topic_id=%s&partition_id=%s&consumer_group=%s", cls.Host, topicID, partitionID, consumerGroup), nil)
	if err != nil {
		return cursors, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return cursors, err
	}
	if resp.StatusCode != 200 {
		return cursors, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &cursors); err != nil {
		return cursors, err
	}

	return cursors, nil
}

func (cls *ClSCleint) GetConsumerData(topicID string, partitionID int, cursor string, count int) (list LogGroupList, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}, "partition_id": {fmt.Sprintf("%s", partitionID)}, "cursor": {fmt.Sprintf("%s", cursor)}, "count": {fmt.Sprintf("%s", count)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/pulllogs", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/pulllogs?topic_id=%s&partition_id=%s&cursor=%s&count=%d", cls.Host, topicID, partitionID, cursor, count), nil)
	if err != nil {
		return list, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return list, err
	}
	if resp.StatusCode != 200 {
		return list, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &list); err != nil {
		return list, err
	}

	return list, nil
}

func (cls *ClSCleint) GetConsumerHeartbeat(heartbeat *ConsumerHeartbeat, topicID string) (partitionIDList []string, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}}
	data, err := json.Marshal(heartbeat)
	if err != nil {
		return partitionIDList, err
	}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/consumerheartbeat", params, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/consumerheartbeat?topic_id=%s", cls.Host, topicID), body)
	if err != nil {
		return partitionIDList, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return partitionIDList, err
	}

	if resp.StatusCode != 200 {
		return partitionIDList, fmt.Errorf("%d", resp.StatusCode)
	}

	list := ConsumerHeartbeat{}
	bod, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bod, &list); err != nil {
		fmt.Println(err)
		return partitionIDList, err
	}

	return list.PartitionIDList, nil
}

func (cls *ClSCleint) GetConsumerGroups(topicID string) (groups ConsumerGroups, err error) {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/consumergroups", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/consumergroups?topic_id=%s", cls.Host, topicID), nil)
	if err != nil {
		return groups, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return groups, err
	}
	if resp.StatusCode != 200 {
		return groups, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &groups); err != nil {
		fmt.Println(err)
		return groups, err
	}

	return groups, nil
}

func (cls *ClSCleint) UpdateConsumerGroup(group *ConsumerGroup, topicID, consumerGroup string) error {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}, "consumer_group": {fmt.Sprintf("%s", consumerGroup)}}
	data, err := json.Marshal(group)
	if err != nil {
		return err
	}

	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/consumergroup", params, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/consumergroup?topic_id=%s&consumer_group=%s", cls.Host, topicID, consumerGroup), body)
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

func (cls *ClSCleint) UpdateConsumerGroupCursor(cursor *ConsumerGroupCursor, topicID, consumerGroup string) error {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}, "consumer_group": {fmt.Sprintf("%s", consumerGroup)}}
	data, err := json.Marshal(cursor)
	if err != nil {
		return err
	}

	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/consumergroupcursor", params, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/consumergroupcursor?topic_id=%s&consumer_group=%s", cls.Host, topicID, consumerGroup), body)
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

func (cls *ClSCleint) DeleteConsumerGroup(topicID, consumerGroup string) error {
	var params = url.Values{"topic_id": {fmt.Sprintf("%s", topicID)}, "consumer_group": {fmt.Sprintf("%s", consumerGroup)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"DELETE", "/consumergroup", params, headers, 300)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/consumergroup?topic_id=%s&consumer_group=%s", cls.Host, topicID, consumerGroup), nil)
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
