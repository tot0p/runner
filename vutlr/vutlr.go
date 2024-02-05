package vutlr

import (
	"io"
	"net/http"
)

type Vutlr struct {
	APIKey string
}

func New() *Vutlr {
	return &Vutlr{}
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	return resp
}
