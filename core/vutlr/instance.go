package vutlr

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Instance struct {
	ID               string   `json:"id"`
	Os               string   `json:"os"`
	Ram              int      `json:"ram"`
	Disk             int      `json:"disk"`
	MainIP           string   `json:"main_ip"`
	VcpuCount        int      `json:"vcpu_count"`
	Region           string   `json:"region"`
	Plan             string   `json:"plan"`
	DateCreated      string   `json:"date_created"`
	Status           string   `json:"status"`
	AllowedBandwidth int      `json:"allowed_bandwidth"`
	NetmaskV4        string   `json:"netmask_v4"`
	GatewayV4        string   `json:"gateway_v4"`
	PowerStatus      string   `json:"power_status"`
	ServerStatus     string   `json:"server_status"`
	V6Network        string   `json:"v6_network"`
	V6MainIP         string   `json:"v6_main_ip"`
	V6NetworkSize    int      `json:"v6_network_size"`
	Label            string   `json:"label"`
	InternalIP       string   `json:"internal_ip"`
	KVM              string   `json:"kvm"`
	Hostname         string   `json:"hostname"`
	OsID             int      `json:"os_id"`
	AppID            int      `json:"app_id"`
	ImageId          string   `json:"image_id"`
	FireWallGroupID  string   `json:"firewall_group_id"`
	Features         []string `json:"features"`
	Tags             []string `json:"tags"`
	UserScheme       string   `json:"user_scheme"`
	DefaultPassword  string   `json:"default_password"`
}

func (i *Instance) String() string {
	return fmt.Sprintf("ID: %s\nOS: %s\nRAM: %d\nDisk: %d\nMainIP: %s\nVcpuCount: %d\nRegion: %s\nPlan: %s\nDateCreated: %s\nStatus: %s\nAllowedBandwidth: %d\nNetmaskV4: %s\nGatewayV4: %s\nPowerStatus: %s\nServerStatus: %s\nV6Network: %s\nV6MainIP: %s\nV6NetworkSize: %d\nLabel: %s\nInternalIP: %s\nKVM: %s\nHostname: %s\nOsID: %d\nAppID: %d\nImageId: %s\nFireWallGroupID: %s\nFeatures: %v\nTags: %v\nUserScheme: %s\nDefaultPassword: %s\n", i.ID, i.Os, i.Ram, i.Disk, i.MainIP, i.VcpuCount, i.Region, i.Plan, i.DateCreated, i.Status, i.AllowedBandwidth, i.NetmaskV4, i.GatewayV4, i.PowerStatus, i.ServerStatus, i.V6Network, i.V6MainIP, i.V6NetworkSize, i.Label, i.InternalIP, i.KVM, i.Hostname, i.OsID, i.AppID, i.ImageId, i.FireWallGroupID, i.Features, i.Tags, i.UserScheme, i.DefaultPassword)
}

// ListIntances = /v2/instances

type ListInstancesResponse struct {
	Instances []Instance `json:"instances"`
	Meta      struct {
		Total int `json:"total"`
		Links struct {
			Next string `json:"next"`
			Prev string `json:"prev"`
		} `json:"links"`
	} `json:"meta"`
}

func (v *Vutlr) ListInstances() ListInstancesResponse {
	resp := v.request(newRequestNoBody(v.rootAPI+"/instances", "GET"))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response ListInstancesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response
}

// CreateInstance = /v2/instances

func (v *Vutlr) CreateInstance(jsonBody string) (Instance, error) {
	bodyR := io.NopCloser(strings.NewReader(jsonBody))
	resp := v.request(newRequest(v.rootAPI+"/instances", "POST", bodyR))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode/100 != 2 {
		return Instance{}, fmt.Errorf("error: %s", string(body))
	}
	var response struct {
		Instance Instance `json:"instance"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response.Instance, nil
}

// GetInstance = /v2/instances/{instance_id}

func (v *Vutlr) GetInstance(instanceID string) (Instance, error) {
	resp := v.request(newRequestNoBody(v.rootAPI+"/instances/"+instanceID, "GET"))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode/100 != 2 {
		return Instance{}, fmt.Errorf("error: %s", string(body))
	}

	var response struct {
		Instance Instance `json:"instance"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	return response.Instance, nil
}

// DeleteInstance = /v2/instances/{instance_id}

func (v *Vutlr) DeleteInstance(instanceID string) error {
	resp := v.request(newRequestNoBody(v.rootAPI+"/instances/"+instanceID, "DELETE"))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("error: %s", string(body))
	}
	return nil
}
