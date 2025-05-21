package model

import "time"

type ShipData struct {
	Data struct {
		Ship
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
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type Origin struct {
	Symbol       string `json:"symbol"`
	Type         string `json:"type"`
	SystemSymbol string `json:"systemSymbol"`
	X            int    `json:"x"`
	Y            int    `json:"y"`
}

type Crew struct {
	Current  int    `json:"current,"`
	Required int    `json:"required"`
	Capacity int    `json:"capacity"`
	Rotation string `json:"rotation"`
	Morale   int    `json:"morale"`
	Wages    int    `json:"wages"`
}

type Frame struct {
	Symbol         string        `json:"symbol"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Condition      int           `json:"condition"`
	Integrity      int           `json:"integrity"`
	ModuleSlots    int           `json:"moduleSlots"`
	MountingPoints int           `json:"mountingPoints"`
	FuelCapacity   int           `json:"fuelCapacity"`
	Requirements   *Requirements `json:"requirements"`
	Quality        int           `json:"quality"`
}

type Requirements struct {
	Power int `json:"power"`
	Crew  int `json:"crew"`
	Slots int `json:"slots"`
}

type Reactor struct {
	Symbol       string        `json:"symbol"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Condition    int           `json:"condition"`
	Integrity    int           `json:"integrity"`
	PowerOutput  int           `json:"powerOutput"`
	Requirements *Requirements `json:"requirements"`
	Quality      int           `json:"quality"`
}

type Engine struct {
	Symbol       string        `json:"symbol"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Condition    int           `json:"condition"`
	Integrity    int           `json:"integrity"`
	Speed        int           `json:"speed"`
	Requirements *Requirements `json:"requirements"`
	Quality      int           `json:"quality"`
}

type Cooldown struct {
	ShipSymbol       string    `json:"shipSymbol"`
	TotalSeconds     int       `json:"total_seconds"`
	RemainingSeconds int       `json:"remainingSeconds"`
	Expiration       time.Time `json:"expiration"`
}

type Modules struct {
	Symbol       string        `json:"symbol"`
	Capacity     int           `json:"capacity"`
	Range        int           `json:"range"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Requirements *Requirements `json:"requirements"`
}

type Mounts struct {
	Symbol       string        `json:"symbol"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Strength     int           `json:"strength"`
	Deposits     []string      `json:"deposits"`
	Requirements *Requirements `json:"requirements"`
}

type Cargo struct {
	Capacity  int         `json:"capacity"`
	Units     int         `json:"units"`
	Inventory []Inventory `json:"inventory"`
}

type Inventory struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Units       int    `json:"units"`
}

type Fuel struct {
	Current  int       `json:"current"`
	Capacity int       `json:"capacity"`
	Consumed *Consumed `json:"consumed"`
}

type Consumed struct {
	Amount    int       `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}
