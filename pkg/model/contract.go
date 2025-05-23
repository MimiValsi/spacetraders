package model

import "time"

type ContractData struct {
	Data struct {
		Contract
	} `json:"data"`
}

type Contract struct {
	Id               string    `json:"id"`
	FactionSymbol    string    `json:"factionSymbol"`
	Type             string    `json:"type"`
	Terms            *Terms    `json:"terms"`
	Accepted         bool      `json:"accepted"`
	Fulfilled        bool      `json:"fulfilled"`
	DeadlineToAccept time.Time `json:"deadlineToAccept"`
}

type Terms struct {
	Deadline time.Time `json:"deadline"`
	Payment  *Payment  `json:"payment"`
	Deliver  []Deliver `json:"deliver"`
}

type Deliver struct {
	TradeSymbol       string `json:"tradeSymbol"`
	DestinationSymbol string `json:"destinationSymbol"`
	UnitsRequired     int32  `json:"unitsRequired"`
	UnitsFulfilled    int32  `json:"unitsFulfilled"`
}

type Payment struct {
	OnAccepted  int32 `json:"onAccepted"`
	OnFulfilled int32 `json:"onFulilled"`
}
