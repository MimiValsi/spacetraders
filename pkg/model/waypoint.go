package model

type WaypointData struct {
	Data struct {
		Waypoint `json:"waypoint"`
	} `json:"data"`
}

type Waypoint struct {
	Symbol              string             `json:"symbol"`
	Type                string             `json:"type"`
	SystemSymbol        string             `json:"systemSymbol"`
	X                   int                `json:"x"`
	Y                   int                `json:"y"`
	Orbitals            []WaypointOrbital  `json:"orbitals"`
	Orbits              string             `json:"orbits,omitempty"`
	Faction             string             `json:"faction,omitempty"`
	Traits              []WaypointTrait    `json:"traits"`
	Modifiers           []WaypointModifier `json:"modifiers,omitempty"`
	Chart               `json:"chart"`
	IsUnderConstruction bool `json:"isUnderConstruction"`
}

type WaypointOrbital struct {
	Symbol string `json:"symbol"`
}

type WaypointModifier struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WaypointTrait struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Chart struct {
	WaypointSymbol string `json:"waypointSymbol"`
	SubmittedBy    string `json:"submittedBy"`
	SubmittedOn    string `json:"submittedOn"`
}
