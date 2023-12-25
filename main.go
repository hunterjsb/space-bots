package main

import (
	"fmt"
	"log"
	"spacebots/models"
)

func main() {
	player := &models.Player

	fleets, err := player.Fleets()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("fleets: %v\n", len(fleets))
	fleet := fleets[0]
	fmt.Printf("%+v\n", fleet.CurrentSystem())
}
