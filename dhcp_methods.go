package provisionclient

import (
	"encoding/json"
	"fmt"
)

type DHCPMethods struct {
	Client *Client
}

//Push Calls

func (dhcp *DHCPMethods) PushPoolByID(poolId string) (*string, error) {

	body, err := dhcp.Client.doRequest("POST", "/dhcp/pools/"+poolId+"/push", nil)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	var pid string
	if val, ok := jsonMap["pid"]; ok {
		pid = fmt.Sprintf("%v", val)
	}

	return &pid, nil
}

func (dns *DHCPMethods) GetPoolPushStatus(poolId, PID string) ([]DHCPPushStatusMessage, error) {
	body, err := dns.Client.doRequest("GET", "/dhcp/pools/"+poolId+"/push_status/"+PID, nil)
	if err != nil {
		return nil, err
	}

	messages := []DHCPPushStatusMessage{}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (dhcp *DHCPMethods) PushGroupByID(groupId string) (*string, error) {

	body, err := dhcp.Client.doRequest("POST", "/dhcp/groups/"+groupId+"/push", nil)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	var pid string
	if val, ok := jsonMap["pid"]; ok {
		pid = fmt.Sprintf("%v", val)
	}

	return &pid, nil
}

func (dns *DHCPMethods) GetGroupPushStatus(groupId, PID string) ([]DHCPPushStatusMessage, error) {
	body, err := dns.Client.doRequest("GET", "/dhcp/groups/"+groupId+"/push_status/"+PID, nil)
	if err != nil {
		return nil, err
	}

	messages := []DHCPPushStatusMessage{}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (dhcp *DHCPMethods) PushServerByID(serverId string) (*string, error) {

	body, err := dhcp.Client.doRequest("POST", "/dhcp/servers/"+serverId+"/push", nil)
	if err != nil {
		return nil, err
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return nil, err
	}

	var pid string
	if val, ok := jsonMap["pid"]; ok {
		pid = fmt.Sprintf("%v", val)
	}

	return &pid, nil
}

func (dns *DHCPMethods) GetServerPushStatus(serverId, PID string) ([]DHCPPushStatusMessage, error) {
	body, err := dns.Client.doRequest("GET", "/dhcp/servers/"+serverId+"/push_status/"+PID, nil)
	if err != nil {
		return nil, err
	}

	messages := []DHCPPushStatusMessage{}
	err = json.Unmarshal(body, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}
