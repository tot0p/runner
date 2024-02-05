package main

import (
	"fmt"
	"github.com/tot0p/env"
	"os"
	"os/signal"
	"runner/vutlr"
	"slices"
	"strings"
	"time"
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

	js, err := os.ReadFile("vm.json")
	if err != nil {
		panic(err)
	}
	i, err := api.CreateInstance(string(js))
	if err != nil {
		panic(err)
	}

	fmt.Println("Creating instance... ", i.ID)
	fmt.Println("Password: ", i.DefaultPassword)

	lastStatus := []string{"none"}
	repass := true
	for i.ServerStatus != "ok" && repass {
		i, err = api.GetInstance(i.ID)
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
		if i.ServerStatus != lastStatus[0] {
			fmt.Println(i.ServerStatus)
			if slices.Contains(lastStatus, i.ServerStatus) {
				repass = false
			}
			lastStatus = append(lastStatus, i.ServerStatus)
		}
	}

	fmt.Println("Instance created")

	// handle ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		err := api.DeleteInstance(i.ID)
		if err != nil {
			panic(err)
		}
		fmt.Println("Instance deleted")
		os.Exit(0)
	}()

	// connect to instance by ssh

	for {
		fmt.Print(">>")
		var cmd string = ""
		_, err := fmt.Scanln(&cmd)
		if err != nil {
			err2 := api.DeleteInstance(i.ID)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println("Instance deleted")
			panic(err)
		}
		cmd = strings.ToLower(cmd)
		switch cmd {
		case "quit":
			err := api.DeleteInstance(i.ID)
			if err != nil {
				panic(err)
			}
			fmt.Println("Instance deleted")
			os.Exit(0)
		default:
			fmt.Println("Unknown command")
		}
	}

	/*
		lst := api.ListInstances()

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
