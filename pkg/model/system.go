package model

type System struct {
	Constellation string `json:"constellation"`
	Symbol        string `json:"symbol"`
	SectorSymbol  string `json:"sectionSymbol"`
	Type          string `json:"type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Waypoints     []SystemWaypoint
	Factions      []SystemFaction
	Name          string `json:"name"`
}

type SystemWaypoint struct {
	Symbol   string `json:"symbol"`
	Type     string `json:"type"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
	Orbitals []WaypointOrbital
	Orbits   string `json:"orbits"`
}

type SystemFaction struct {
	Symbol string `json:"symbol"`
}
