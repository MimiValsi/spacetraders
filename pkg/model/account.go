package model

import "time"

type AccountData struct {
	Data struct {
		Account Account
	} `json:"data"`
}

type Account struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}
