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
	/*
	   	i, err := api.CreateInstance(`{
	     "region": "ewr",
	     "plan": "vc2-1c-2gb",
	     "label": "ApiRunner",
	     "os_id": 477,
	     "hostname": "runner"
	   }
	   			`)
	   	if err != nil {
	   		panic(err)
	   	}
	   	fmt.Println(i.ID)

	*/

	lst := api.ListInstances()
	fmt.Println(lst.Meta.Total)

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
}
