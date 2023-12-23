package models

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Key        string // Not part of JSON, used internally
	Credits    int    `json:"credits"`
	CreatedAt  string `json:"createdAt"`
	Registered bool   `json:"registered"`
}

func (u *User) Me() {
	e := Endpoint{"/users/me", "GET"}
	resp, err := e.Request(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(body, u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", u)
}
