package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Station struct {
	DirectSell bool `json:"directSell"`
	BuyShips   bool `json:"buyShips"`
}

type System struct {
	ID                 string              `json:"id"`
	Name               string              `json:"name"`
	Asteroid           map[string]string   `json:"asteroid"`
	Station            Station             `json:"station"`
	X                  int                 `json:"x"`
	Y                  int                 `json:"y"`
	NeighboringSystems []map[string]string `json:"neighboringSystems"`
}

func GetSystem(systemID string) (*System, error) {
	e := Endpoint{"/systems/" + systemID, "GET"}
	resp, err := e.Request(nil)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var s System
	json.Unmarshal(body, &s)
	return &s, nil
}

func (s *System) About() error {
	if s.ID == "" {
		return errors.New("System is missing ID")
	}

	e := Endpoint{"/systems/" + s.ID, "GET"}
	resp, err := e.Request(nil)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, s)
	if err != nil {
		return err
	}

	fmt.Printf("system: %v\n", s)
	return nil
}
