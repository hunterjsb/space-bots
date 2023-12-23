package models

import (
	"encoding/json"
	"fmt"
	"io"
)

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Key        string // Not part of JSON, used internally
	Credits    int    `json:"credits"`
	CreatedAt  string `json:"createdAt"`
	Registered bool   `json:"registered"`
}

var endpoints map[string]*Endpoint = map[string]*Endpoint{
	"Me":     {"/users/me", "GET"},
	"Fleets": {"/fleets/my", "GET"},
}

func (u *User) Me() error {
	e := endpoints["Me"]
	resp, err := e.Request(nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, u)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", u)
	return nil
}

func (u *User) Fleets() ([]Fleet, error) {
	e := &Endpoint{"/fleets/my", "GET"}
	resp, err := e.Request(nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var fleets []Fleet
	err = json.Unmarshal(body, &fleets)
	if err != nil {
		return nil, err
	}

	return fleets, nil
}
