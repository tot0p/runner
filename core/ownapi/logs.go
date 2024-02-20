package ownapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runner/core/vutlr"
)

type Item struct {
	containerID     string `bson:"containerID"`
	repositoryURL   string `bson:"repositoryURL"`
	creationDate    string `bson:"creationDate"`
	destructionDate string `bson:"destructionDate"`
}

func (i Item) String() string {
	return fmt.Sprintf("ContainerID: %s\nRepositoryURL: %s\nCreationDate: %s\nDestructionDate: %s\n", i.containerID, i.repositoryURL, i.creationDate, i.destructionDate)
}

func Logs(cmd []string, i *vutlr.Instance) {
	req, err := http.NewRequest("GET", "http://"+i.MainIP+":80/logs", nil)
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
		var result []Item
		err := json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, v := range result {
			fmt.Println(v)
		}
	}
}
