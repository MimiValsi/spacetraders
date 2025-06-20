// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"time"
)

type Account struct {
	ID        string
	Email     string
	CreatedAt time.Time
}

type Agent struct {
	ID              int32
	AccountID       string
	Token           string
	Symbol          string
	Headquarters    string
	Credits         int64
	StartingFaction string
	ShipCount       int32
}

type Cargo struct {
	ID       int32
	Capacity int32
	Units    int32
	ShipID   int32
}

type Consumed struct {
	ID       int32
	Amount   int32
	Timestmp time.Time
	FuelID   int32
}

type Contract struct {
	Cid              int32
	ID               string
	FactionSymbol    string
	Type             string
	Accepted         bool
	Fulfilled        bool
	DeadlineToAccept time.Time
	AgentID          int32
}

type Cooldown struct {
	ID               int32
	ShipSymbol       string
	TotalSeconds     int32
	RemainingSeconds int32
	Expiration       time.Time
	ShipID           int32
}

type Crew struct {
	ID       int32
	Current  int32
	Required int32
	Capacity int32
	Rotation string
	Morale   int32
	Wages    int32
	ShipID   int32
}

type Deliver struct {
	ID                int32
	TradeSymbol       string
	DestinationSymbol string
	UnitsRequired     int32
	UnitsFulfilled    int32
	TermID            int32
}

type Destination struct {
	ID           int32
	Symbol       string
	Type         string
	SystemSymbol string
	X            int32
	Y            int32
	RouteID      int32
}

type Engine struct {
	ID            int32
	Symbol        string
	Name          string
	Condition     float32
	Integrity     float32
	Description   string
	Speed         int32
	Quality       int32
	RequirementID int32
	ShipID        int32
}

type Faction struct {
	ID           int32
	Symbol       string
	Name         string
	Description  string
	Headquarters string
	IsRecruiting bool
	AgentID      int32
}

type Frame struct {
	ID             int32
	Symbol         string
	Name           string
	Condition      float32
	Integrity      float32
	Description    string
	ModuleSlots    int32
	MountingPoints int32
	FuelCapacity   int32
	Quality        int32
	RequirementID  int32
	ShipID         int32
}

type Fuel struct {
	ID       int32
	Current  int32
	Capacity int32
	ShipID   int32
}

type Inventory struct {
	ID          int32
	Symbol      string
	Name        string
	Description string
	Units       int32
	CargoID     int32
}

type Module struct {
	ID            int32
	Symbol        string
	Name          string
	Description   string
	Capacity      int32
	Range         int32
	RequirementID int32
	ShipID        int32
}

type Mount struct {
	ID            int32
	Symbol        string
	Name          string
	Description   string
	Strength      int32
	Deposits      []string
	RequirementID int32
	ShipID        int32
}

type Nav struct {
	ID             int32
	SystemSymbol   string
	WaypointSymbol string
	Status         string
	FlightMode     string
	ShipID         int32
}

type Origin struct {
	ID           int32
	Symbol       string
	Type         string
	SystemSymbol string
	X            int32
	Y            int32
	RouteID      int32
}

type Payment struct {
	ID          int32
	OnAccepted  int32
	OnFulfilled int32
	TermID      int32
}

type Reactor struct {
	ID            int32
	Symbol        string
	Name          string
	Condition     float32
	Integrity     float32
	Description   string
	PowerOutput   int32
	Quality       int32
	RequirementID int32
	ShipID        int32
}

type Registration struct {
	ID            int32
	Name          string
	FactionSymbol string
	Role          string
	ShipID        int32
}

type Requirement struct {
	ID    int32
	Power int32
	Crew  int32
	Slots int32
}

type Route struct {
	ID            int32
	DepartureTime time.Time
	Arrival       time.Time
	NavID         int32
}

type Ship struct {
	ID      int32
	Symbol  string
	AgentID int32
}

type Term struct {
	ID         int32
	Deadline   time.Time
	ContractID int32
}

type Token struct {
	ID string
}

type Trait struct {
	ID          int32
	Symbol      string
	Name        string
	Description string
	FactionID   int32
}
