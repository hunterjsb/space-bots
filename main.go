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
	sys := fleet.CurrentSystem()
	fmt.Println("neighbors: ", sys.NeighboringSystems)

	// lastSys := &models.System{ID: sys.NeighboringSystems[len(sys.NeighboringSystems)-1]["systemId"]}
	// lastSys.About()

	// fleet.Travel(lastSys)
	// fleet.CurrentSystem()
}
