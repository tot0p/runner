package main

import (
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
	vultr := vutlr.New()
	vultr.SetAPIKey(env.Get("API_KEY"))
}
