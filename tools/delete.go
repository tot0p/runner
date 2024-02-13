package main

import (
	"github.com/tot0p/env"
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

	err := api.DeleteInstance("8683c4aa-9bea-4018-bc87-937e94d05695")
	if err != nil {
		panic(err)
	}
}
