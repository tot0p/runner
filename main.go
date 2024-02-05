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
	api := vutlr.New()
	api.SetAPIKey(env.Get("API_KEY"))

	/*
		js, err := os.ReadFile("vm.json")
		if err != nil {
			panic(err)
		}
		i, err := api.CreateInstance(string(js))
		if err != nil {
			panic(err)
		}
		fmt.Println(i.ID)

		lst := api.ListInstances()
		fmt.Println(lst.Meta.Total)

	*/

	/*
		id := lst.Instances[0].ID

		fmt.Println(id)

		i, err := api.GetInstance(id)
		if err != nil {
			panic(err)
		}

		fmt.Println(i.ID)
		fmt.Println(i.Os)

	*/
	/*
		for i.ServerStatus != "ok" {
			i, err = api.GetInstance(i.ID)
			if err != nil {
				panic(err)
			}
			fmt.Println(i.ServerStatus)
		}

		if i.Label == "ApiRunner" {
			fmt.Println("Instance created")
			err := api.DeleteInstance(i.ID)
			if err != nil {
				panic(err)
			}
			fmt.Println("Instance deleted")
		}

	*/
}
