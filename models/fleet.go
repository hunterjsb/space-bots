package models

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
