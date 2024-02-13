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

	fmt.Println("API : ", i.MainIP+":8080")
	// run commands

	cmd := "apt update && apt install -y git && git clone https://github.com/tot0p/api_runner && curl -LO https://go.dev/dl/go1.22.0.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz && export PATH=$PATH:/usr/local/go/bin && cd api_runner && go build . && iptables -A INPUT -p tcp --dport 8080 -j ACCEPT && ufw allow 8080/tcp && ./api_runner"

	session, err := client.NewSession()
	if err != nil {
		core.Close(api, i)
		panic(err)
	}

	/*
		ss, err := session.StdoutPipe()
		if err != nil {
			core.Close(api, i)
			panic(err)
		}
		_, err = io.Copy(os.Stdout, ss)
		if err != nil {
			core.Close(api, i)
			return
		}

	*/

	fmt.Println("Running command: ", cmd)
	out, err := session.CombinedOutput(cmd)
	if err != nil {
		fmt.Println("Error running command: ", cmd)
		fmt.Println(string(out))
	} else {
		fmt.Println(string(out))
	}

	defer func() {
		err := session.Close()
		if err != nil {
			core.Close(api, i)
			panic(err)
		}
	}()

	for {
		fmt.Print(">>")
		var cmd string = ""
		_, err := fmt.Scanln(&cmd)
		if err != nil {
			core.Close(api, i)
			panic(err)
		}
		cmd = strings.ToLower(cmd)
		switch cmd {
		case "quit":
			core.Close(api, i)
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
