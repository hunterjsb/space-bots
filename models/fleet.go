package models

import (
	"bytes"
	"encoding/json"
	"errors"
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

type ActionResult struct {
	Duration   float64 `json:"duration"`
	FinishTime string  `json:"finishTime"`
}

func (f *Fleet) Get() error {
	e := Endpoint{"/fleets/" + f.ID, "GET"}
	resp, err := e.Request(nil)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	json.Unmarshal(body, f)
	return nil
}

func (f *Fleet) System() *System {
	sys, err := GetSystem(f.LocationSystemId)
	if err != nil {
		log.Fatal(err)
	}
	return sys
}

func (f *Fleet) Mine() (*ActionResult, error) {
	e := &Endpoint{"/fleets/" + f.ID + "/mine", "POST"}
	resp, err := e.Request(nil)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res ActionResult
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (f *Fleet) DirectSell(resouces map[string]int) (int, error) {
	e := Endpoint{"/fleets/" + f.ID + "/direct-sell", "POST"}
	jsonData, err := json.Marshal(map[string]map[string]int{"resources": resouces})
	if err != nil {
		return 0, err
	}
	resp, err := e.Request(bytes.NewReader(jsonData))
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New("Error selling resources, " + fmt.Sprint(resp.StatusCode))
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var data map[string]int
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}
	return data["creditsGained"], nil
}

func (f *Fleet) Travel(sys *System) (*ActionResult, error) {
	// I don't think this works
	jsonData, err := json.Marshal(map[string]string{"destinationSystemId": sys.ID})
	if err != nil {
		return nil, err
	}

	e := &Endpoint{"/fleets/" + f.ID + "/travel", "POST"}
	resp, err := e.Request(bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res ActionResult
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("Error travelling to system, " + fmt.Sprint(resp.StatusCode))
	}
	return &res, nil
}

func (f *Fleet) Buy(ships map[string]int) (int, error) {
	jsonData, err := json.Marshal(map[string]map[string]int{"shipsToBuy": ships})
	if err != nil {
		return 0, err
	}
	e := Endpoint{"/fleets/" + f.ID + "/buy-ships", "POST"}
	resp, err := e.Request(bytes.NewReader(jsonData))
	if err != nil {
		return 0, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	var creditsSpent map[string]int
	json.Unmarshal(body, &creditsSpent)
	fmt.Println("BUY RESULTS: ", resp.StatusCode, creditsSpent)
	return creditsSpent["creditsSpent"], nil
}

func ShipTypes() error {
	e := Endpoint{"/ship-types", "GET"}
	resp, err := e.Request(nil)
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
