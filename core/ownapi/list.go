package ownapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runner/core/vutlr"
)

type Vm struct {
	ID     string `json:"Id"`
	IP     string `json:"IPAddress"`
	Name   string `json:"Names"`
	State  string `json:"State"`
	Status string `json:"Status"`
	Image  string `json:"Image"`
}

type ListVm []Vm

func (l ListVm) String() string {
	var result string
	for _, v := range l {
		result += fmt.Sprintf("ID: %s\nIP: %s\nName: %s\nState: %s\nStatus: %s\nImage: %s\n\n", v.ID, v.IP, v.Name, v.State, v.Status, v.Image)
	}
	return result
}

func List(cmd []string, i *vutlr.Instance) {
	req, err := http.NewRequest("GET", "http://"+i.MainIP+":80/vm", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp, err := http.DefaultClient.Do(req)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.Body != nil {
		var body []byte
		body, err = io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		var result ListVm
		err := json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)
	}
}
