package core

import (
	"fmt"
	"github.com/tot0p/env"
	"os"
	"runner/core/vutlr"
	"time"
)

// Connect creates a new instance and returns the API, password and instance
func Connect() (*vutlr.Vutlr, string, vutlr.Instance) {
	api := vutlr.New()
	api.SetAPIKey(env.Get("API_KEY"))

	js, err := os.ReadFile("vm.json")
	if err != nil {
		panic(err)
	}
	i, err := api.CreateInstance(string(js))
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating instance... ", i.ID)
	fmt.Println("IP: ", i.MainIP)
	fmt.Println("Password: ", i.DefaultPassword)
	pass := i.DefaultPassword

	lastStatus := "none"
	count := 0
	for count < 60 {
		i, err = api.GetInstance(i.ID)
		if err != nil {
			panic(err)
		}

		time.Sleep(1 * time.Second)
		count++
		if i.ServerStatus != lastStatus {
			fmt.Println(i.ServerStatus)
			lastStatus = i.ServerStatus
		}
	}

	return api, pass, i
}

// Close deletes the instance
func Close(api *vutlr.Vutlr, i vutlr.Instance) {
	err := api.DeleteInstance(i.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Instance deleted")
}
