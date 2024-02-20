package main

import (
	"errors"
	"fmt"
	"github.com/tot0p/env"
	"golang.org/x/crypto/ssh"
	"io"
	"net/http"
	"os"
	"os/signal"
	"runner/core"
	"runner/core/ownapi"
	"strings"
	"time"
)

type newline struct{ tok string }

func (n *newline) Scan(state fmt.ScanState, verb rune) error {
	tok, err := state.Token(false, func(r rune) bool {
		return r != '\n'
	})
	if err != nil {
		return err
	}
	if _, _, err := state.ReadRune(); err != nil {
		if len(tok) == 0 {
			panic(err)
		}
	}
	n.tok = string(tok)
	return nil
}

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

	// Try to connect 5 times
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
	go func() {
		cmd := "apt update && apt install -y git && git clone https://github.com/tot0p/api_runner && curl -fsSL https://get.docker.com/ -o install-docker.sh && sh install-docker.sh && curl -LO https://go.dev/dl/go1.22.0.linux-amd64.tar.gz && rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz && export PATH=$PATH:/usr/local/go/bin && apt-get install -y gnupg && wget -qO - https://www.mongodb.org/static/pgp/server-6.0.asc | apt-key add -  && echo \"deb http://repo.mongodb.org/apt/debian buster/mongodb-org/6.0 main\" | sudo tee /etc/apt/sources.list.d/mongodb-org-6.0.list && apt-get update && apt-get install -y mongodb-org && systemctl start mongod && cd api_runner && go build . && ./api_runner"

		session, err := client.NewSession()
		if err != nil {
			core.Close(api, i)
			panic(err)
		}

		fmt.Println("Running command: ", cmd)
		out, err := session.CombinedOutput(cmd) // with the script there are no output because running infinite loop
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
	}()

	deploy := false

	req, err := http.NewRequest("GET", "http://"+i.MainIP+":80/ping", nil)
	if err != nil {
		core.Close(api, i)
		panic(err)
	}

	// handle if /ping is available
	for !deploy {
		// send request to api

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			if strings.Contains(err.Error(), "dial tcp") {
				time.Sleep(2 * time.Second)
				continue
			}
			core.Close(api, i)
			panic(err)
		}
		if resp.Body != nil {
			var body []byte
			body, err = io.ReadAll(resp.Body)

			if string(body) == "{\n    \"message\": \"pong\"\n}" {
				fmt.Println("Deployed")
				deploy = true
			}
		}
		err = resp.Body.Close()
		if err != nil {
			core.Close(api, i)
			panic(err)
		}
	}

	for {
		fmt.Print(">>")
		var temp newline
		_, err := fmt.Scan(&temp)
		if err != nil {
			core.Close(api, i)
			panic(err)
		}
		temp.tok = strings.TrimSpace(temp.tok)
		temp.tok = strings.ReplaceAll(temp.tok, "\n", "")
		temp.tok = strings.ToLower(temp.tok)
		cmd := strings.Split(temp.tok, " ")
		switch cmd[0] {
		case "clone":
			if len(cmd) != 2 {
				fmt.Println("Usage: clone <url>")
				continue
			}
			// run commands
			ownapi.Clone(cmd, &i)
		case "list":
			ownapi.List(cmd, &i)
		case "close":
			ownapi.Close(cmd, &i)
		case "quit", "exit":
			core.Close(api, i)
			os.Exit(0)
		case "help":
			fmt.Println("clone <url> - clone a repository")
			fmt.Println("list - list running containers")
			fmt.Println("close <instance_id> - close a container")
			fmt.Println("quit, exit - close the program")
		case "logs":
			ownapi.Logs(cmd, &i)
		case "ip":
			fmt.Println(i.MainIP)
		case "pass":
			fmt.Println(pass)
		default:
			fmt.Println("Unknown command")
		}
	}
}
