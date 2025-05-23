package model

import "time"

type ShipData struct {
	Data struct {
		Ship Ship
	} `json:"data"`
}

type Ship struct {
	Symbol       string        `json:"symbol"`
	Registration *Registration `json:"registration"`
	Nav          *Nav          `json:"nav"`
	Crew         *Crew         `json:"crew"`
	Frame        *Frame        `json:"frame"`
	Reactor      *Reactor      `json:"reactor"`
	Engine       *Engine       `json:"engine"`
	Cooldown     *Cooldown     `json:"cooldown"`
	Modules      []Modules     `json:"modules"`
	Mounts       []Mounts      `json:"mounts"`
	Cargo        *Cargo        `json:"cargo"`
	Fuel         *Fuel         `json:"fuel"`
}

type Registration struct {
	Name          string `json:"name"`
	FactionSymbol string `json:"factionSymbol"`
	Role          string `json:"role"`
}

type Nav struct {
	SystemSymbol   string `json:"systemSymbol"`
	WaypointSymbol string `json:"waypointSymbol"`
	Route          *Route `json:"route"`
	Status         string `json:"status"`
	FlightMode     string `json:"flightMode"`
}

type Route struct {
	Destination   *Destination `json:"destination"`
	Origin        *Origin      `json:"origin"`
	DepartureTime time.Time    `json:"departureTime"`
	Arrival       time.Time    `json:"arrival"`
}

type Destination struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int32  `json:"x"`
	Y            int32  `json:"y"`
}

type Origin struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int32  `json:"x"`
	Y            int32  `json:"y"`
}

type Crew struct {
	Current  int32  `json:"current,"`
	Required int32  `json:"required"`
	Capacity int32  `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int32  `json:"morale"`
	Wages    int32  `json:"wages"`
}

type Frame struct {
	Symbol         string        `json:"symbol"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Condition      float32       `json:"condition"`
	Integrity      float32       `json:"integrity"`
	ModuleSlots    int32         `json:"moduleSlots"`
	MountingPoints int32         `json:"mountingPoints"`
	FuelCapacity   int32         `json:"fuelCapacity"`
	Requirements   *Requirements `json:"requirements"`
	Quality        int32         `json:"quality"`
}

type Requirements struct {
	Power int32 `json:"power"`
	Crew  int32 `json:"crew"`
	Slots int32 `json:"slots"`
}

type Reactor struct {
	Symbol       string        `json:"symbol"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Condition    float32       `json:"condition"`
	Integrity    float32       `json:"integrity"`
	PowerOutput  int32         `json:"powerOutput"`
	Requirements *Requirements `json:"requirements"`
	Quality      int32         `json:"quality"`
}

type Engine struct {
	Symbol       string        `json:"symbol"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Condition    float32       `json:"condition"`
	Integrity    float32       `json:"integrity"`
	Speed        int32         `json:"speed"`
	Requirements *Requirements `json:"requirements"`
	Quality      int32         `json:"quality"`
}

type Cooldown struct {
	ShipSymbol       string    `json:"shipSymbol"`
	TotalSeconds     int32     `json:"total_seconds"`
	RemainingSeconds int32     `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
}

type Modules struct {
	Symbol       string        `json:"symbol"`
	Capacity     int32         `json:"capacity"`
	Range        int32         `json:"range"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Requirements *Requirements `json:"requirements"`
}

type Mounts struct {
	Symbol       string        `json:"symbol"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Strength     int32         `json:"strength,omitempty"`
	Deposits     []string      `json:"deposits,omitempty"`
	Requirements *Requirements `json:"requirements"`
}

type Cargo struct {
	Capacity  int32       `json:"capacity"`
	Units     int32       `json:"units"`
	Inventory []Inventory `json:"inventory"`
}

type Inventory struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int32  `json:"units"`
}

type Fuel struct {
	Current  int32     `json:"current"`
	Capacity int32     `json:"capacity"`
	Consumed *Consumed `json:"consumed"`
}

type Consumed struct {
	Amount    int32     `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
