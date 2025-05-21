package model

type FactionData struct {
	Data struct {
		Faction
	} `json:"data"`
}

type Faction struct {
	Symbol       string   `json:"symbol"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Headquarters string   `json:"headquartes"`
	Traits       []Traits `json:"traits"`
	IsRecruiting bool     `json:"isRecruiting"`
}

type Traits struct {
	Symbol      string `json:"symbol"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
