package cls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Machines struct {
	Machines []Machine `json:"machines"`
}

type Machine struct {
	IP     string `json:"ip"`
	Status string `json:"status"`
}

type MachineGroups struct {
	Groups []MachineGroup `json:"machine_groups"`
}

type MachineGroup struct {
	GroupID    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	Type       string `json:"type"`
	IPs        string `json:"ips"`
	Labels     int    `json:"labels"`
	CreateTime int    `json:"create_time"`
}

// TODO
func (cls *ClSCleint) CreateMachineGroup(machineGroup *MachineGroup) (string, error) {
	data, err := json.Marshal(machineGroup)
	if err != nil {
		return "", err
	}
	fmt.Println(string(data))
	// data := fmt.Sprintf("{\"logset_name\":\"%s\",\"period\":%d}", logSetName, period)
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"POST", "/machinegroup", nil, headers, 300)

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/machinegroup", cls.Host), body)
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

	group := MachineGroup{}
	bod, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(bod, &group); err != nil {
		return "", err
	}

	return group.GroupID, nil
}

func (cls *ClSCleint) GetMachineGroup(groupID string) (group MachineGroup, err error) {
	var params = url.Values{"group_id": {fmt.Sprintf("%s", groupID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/machinegroup", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/machinegroup?group_id=%s", cls.Host, groupID), nil)
	if err != nil {
		return group, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return group, err
	}
	if resp.StatusCode != 200 {
		return group, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &group); err != nil {
		fmt.Println(err)
		return group, err
	}

	return group, nil
}

func (cls *ClSCleint) GetMachines(groupID string) (machines Machines, err error) {
	var params = url.Values{"group_id": {fmt.Sprintf("%s", groupID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/machines", params, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/machines?group_id=%s", cls.Host, groupID), nil)
	if err != nil {
		return machines, err
	}
	req.Header.Add("Authorization", sig)
	req.Header.Add("Host", fmt.Sprintf("%s", cls.Host))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return machines, err
	}
	if resp.StatusCode != 200 {
		return machines, fmt.Errorf("%d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &machines); err != nil {
		fmt.Println(err)
		return machines, err
	}

	return machines, nil
}

func (cls *ClSCleint) GetMachineGroups() (groups MachineGroups, err error) {
	// var params = url.Values{"shipper_id": {fmt.Sprintf("%s", shipperID)}, "start_time": {fmt.Sprintf("%s", startTime)}, "end_time": {fmt.Sprintf("%s", endTime)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"GET", "/machinegroups", nil, headers, 300)
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s/machinegroups", cls.Host), nil)
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

func (cls *ClSCleint) UpdateMachineGroup(group *MachineGroup) error {
	data, err := json.Marshal(group)
	if err != nil {
		return err
	}

	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	body := bytes.NewBuffer([]byte(data))
	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"PUT", "/machinegroup", nil, headers, 300)

	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s/machinegroup", cls.Host), body)
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

func (cls *ClSCleint) DeleteMachineGroup(groupID string) error {
	var params = url.Values{"group_id": {fmt.Sprintf("%s", groupID)}}
	var headers = url.Values{"Host": {fmt.Sprintf("%s", cls.Host)}, "User-Agent": {"AuthSDK"}}

	sig := Signature(fmt.Sprintf("%s", cls.SecretId), fmt.Sprintf("%s", cls.SecretKey),
		"DELETE", "/machinegroup", params, headers, 300)
	req, err := http.NewRequest("DELETE", fmt.Sprintf("http://%s/machinegroup?group_id=%s", cls.Host, groupID), nil)
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
