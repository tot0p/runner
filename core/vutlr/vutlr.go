package vutlr

import (
	"io"
	"net/http"
)

type Vutlr struct {
	APIKey  string
	rootAPI string
}

func New() *Vutlr {
	return &Vutlr{
		rootAPI: "https://api.vultr.com/v2",
	}
}

func (v *Vutlr) SetAPIKey(key string) {
	v.APIKey = key
}

type request struct {
	URL    string
	method string
	body   io.Reader
}

func newRequest(url string, method string, body io.Reader) request {
	return request{
		URL:    url,
		method: method,
		body:   body,
	}
}

func newRequestNoBody(url string, method string) request {
	return request{
		URL:    url,
		method: method,
		body:   nil,
	}
}

// request sends a request to the API with the Bearer token
func (v *Vutlr) request(r request) *http.Response {
	var bearer = "Bearer " + v.APIKey

	req, err := http.NewRequest(r.method, r.URL, r.body)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp
}
