package models

import (
	"io"
	"net/http"
	"spacebots/utils"
)

var client *http.Client = &http.Client{}

const base string = "https://space-bots.longwelwind.net/v1"

var Player User = User{Key: utils.Env["TOKEN"]}

type Endpoint struct {
	Route  string
	Method string // "GET" or "POST"
}

func (e *Endpoint) Request(body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(e.Method, base+e.Route, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+Player.Key)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return client.Do(req)
}
