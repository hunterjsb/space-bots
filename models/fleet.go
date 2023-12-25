package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Fleet struct {
	Cargo            map[string]int    `json:"cargo"`
	CurrentAction    interface{}       `json:"currentAction"`
	ID               string            `json:"id"`
	LocationSystemId string            `json:"locationSystemId"`
	Owner            map[string]string `json:"owner"`
	Ships            map[string]int    `json:"ships"`
}

func (f *Fleet) CurrentSystem() *System {
	sys, err := GetSystem(f.LocationSystemId)
	if err != nil { // I don't actually think this can ever err
		log.Fatal(err)
	}
	return sys
}

func (f *Fleet) Mine() {
	e := &Endpoint{"/fleets/" + f.ID + "/mine", "POST"}
	resp, err := e.Request(nil)
	if err != nil {
		log.Fatal(err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode, data)
}

func (f *Fleet) Travel(sys *System) error {
	// I don't think this works
	jsonData, err := json.Marshal(map[string]string{"destinationSystemId": sys.ID})
	if err != nil {
		return err
	}

	e := &Endpoint{"/fleets/" + f.ID + "/travel", "POST"}
	resp, err := e.Request(bytes.NewReader(jsonData))
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var data interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode, data)
	return nil
}
