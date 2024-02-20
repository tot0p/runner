package ownapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runner/core/utils"
	"runner/core/vutlr"
	"strings"
)

type CloneResponse struct {
	Id      string         `json:"id"`
	Message string         `json:"message"`
	Ports   map[string]int `json:"ports"`
}

func Clone(cmd []string, i *vutlr.Instance) {
	body := strings.NewReader(fmt.Sprintf("{\"link\":\"%s\"}", cmd[1]))
	req, err := http.NewRequest("POST", "http://"+i.MainIP+":80/vm", body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	req.Header.Set("Content-Type", "application/json")
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
		var result CloneResponse
		err := json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err)
			return
		}

		for k, v := range result.Ports {
			if k == "80" || k == "443" {
				utils.Openbrowser(fmt.Sprintf("http://%s:%d", i.MainIP, v))
				fmt.Println("Opening browser on port ", v)
				return
			}
		}
		fmt.Println(string(body))
	}

}
