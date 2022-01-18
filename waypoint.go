package main

type Waypoint struct {
	Coords  []float64 `json:"coords"`
	NavType string    `json:"type"`
	Name    string    `json:"name"`
}
