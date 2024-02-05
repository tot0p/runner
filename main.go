package main

import "github.com/tot0p/env"

func init() {
	err := env.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

}
