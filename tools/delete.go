package main

import (
	"fmt"
	"github.com/tot0p/env"
	"os"
	"runner/core/vutlr"
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

	if len(os.Args) != 2 {
		fmt.Println("Usage: delete <instance_id>")
		os.Exit(1)
	}

	err := api.DeleteInstance(os.Args[1])
	if err != nil {
		panic(err)
	}
}
