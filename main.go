package main

import (
	"fmt"
	"spacebots/models"
	"time"
)

func loopMine(f *models.Fleet) {
	for {
		res, _ := f.Mine()
		fmt.Printf("cargo: %+v current action: %+v action results: %+v\n", f.Cargo, f.CurrentAction, res)
		time.Sleep(time.Duration(res.Duration) * time.Second)
		f.Get()
	}
}

func main() {
	player := &models.Player
	fleets, _ := player.Fleets()

	fmt.Printf("fleets: %v\n", len(fleets))
	fleet := fleets[0]
	loopMine(&fleet)
}
