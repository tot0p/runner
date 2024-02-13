package main

import (
	"errors"
	"fmt"
	"github.com/tot0p/env"
	"golang.org/x/crypto/ssh"
	"os"
	"os/signal"
	"runner/core"
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
	api, pass, i := core.Connect()

	fmt.Println("Instance created")

	// handle ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		core.Close(api, i)
		os.Exit(0)
	}()

	// connect to instance by ssh
	host := i.MainIP + ":22"
	user := "root"
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	var client *ssh.Client = nil
	err := errors.New("not nil")
	count := 0

	for err != nil && count < 5 {
		client, err = ssh.Dial("tcp", host, config)
		if err != nil {
			fmt.Println("Retrying...")
			time.Sleep(2 * time.Second)
			count++
		}
	}

	if client == nil {
		core.Close(api, i)
		panic(err)
	}
	fmt.Println("Connected to instance")

	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	// run commands
	cmds := []string{
		"echo 'hello world';",
		"apt update",
		"apt install -y git",
	}

	session, err := client.NewSession()
	if err != nil {
		core.Close(api, i)
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println("Running command: ", cmd)
		out, err := session.CombinedOutput(cmd)
		if err != nil {
			fmt.Println("Error running command: ", cmd)
			fmt.Println(string(out))
		} else {
			fmt.Println(string(out))
		}

		err = session.Close()
		if err != nil {
			if err != nil {
				err2 := api.DeleteInstance(i.ID)
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("Instance deleted")
			}
			panic(err)
		}
		session, err = client.NewSession()
		if err != nil {
			if err != nil {
				err2 := api.DeleteInstance(i.ID)
				if err2 != nil {
					panic(err2)
				}
				fmt.Println("Instance deleted")
			}
			panic(err)
		}

		time.Sleep(5 * time.Second)

	}

	err = session.Close()
	if err != nil {
		if err != nil {
			err2 := api.DeleteInstance(i.ID)
			if err2 != nil {
				panic(err2)
			}
			fmt.Println("Instance deleted")
		}
		panic(err)
	}

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
