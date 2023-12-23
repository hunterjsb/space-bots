package models

import (
	"bytes"
	"encoding/json"
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
	sys := &System{ID: f.LocationSystemId}
	sys.About()
	return sys
}

func (f *Fleet) Travel(sys *System) {
	jsonData, err := json.Marshal(map[string]string{"destinationSystemId": sys.ID})
	if err != nil {
		log.Fatal(err)
	}

	e := &Endpoint{"/fleets/" + f.ID + "/travel", "POST"}
	_, err = e.Request(bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
}
