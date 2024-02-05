package vutlr

import (
	"io"
)

func (v *Vutlr) ListInstances() {
	resp := v.request(newRequestNoBody(v.rootAPI+"/instances", "GET"))
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
