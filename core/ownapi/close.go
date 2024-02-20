package ownapi

import (
	"fmt"
	"io"
	"net/http"
	"runner/core/vutlr"
)

func Close(cmd []string, i *vutlr.Instance) {
	if len(cmd) != 2 {
		fmt.Println("Usage: close <instance_id>")
		return
	}

	req, err := http.NewRequest("DELETE", "http://"+i.MainIP+":80/vm/"+cmd[1], nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp, err := http.DefaultClient.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
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
		fmt.Println(string(body))
	}

}
