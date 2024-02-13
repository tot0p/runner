package main

import (
	"fmt"
	"github.com/tot0p/env"
	"runner/vutlr"
)

func init() {
	err := env.Load()
	if err != nil {
		panic(err)
	}

}

func main() {
	api := vutlr.New()
	api.SetAPIKey(env.Get("API_KEY"))

	lst := api.ListInstances()

	for _, i := range lst.Instances {
		fmt.Println(i)
	}
}
