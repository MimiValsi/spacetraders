package model

type AgentRegister struct {
	Data struct {
		Token    string   `json:"token"`
		Agent    Agent    `json:"agent"`
		Faction  Faction  `json:"faction"`
		Contract Contract `json:"contract"`
		Ships    []Ship   `json:"ships"`
	} `json:"data"`
}

type AgentData struct {
	Data struct {
		Agent
	} `json:"data"`
}

type Agent struct {
	AccountID       string `json:"accountId,omitempty"`
	Symbol          string `json:"symbol"`
	Headquarters    string `json:"headquarters"`
	Credits         int64  `json:"credits"`
	StartingFaction string `json:"startingFaction"`
	ShipCount       int32  `json:"shipCount"`
}
