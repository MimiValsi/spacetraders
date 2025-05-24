package api

import (
	"context"
	"github.com/MimiValsi/spacetraders/internal/database"
)

type Agent struct {
	Token   string
	Credits int64
	Headquarters string

	DB *database.Queries
}

func LoadAgent(db *database.Queries) (*Agent, error) {
	agent, err := db.GetAgent(context.Background())
	if err != nil {
		return nil, err
	}

	return &Agent{
		Token: agent.Token,
		Credits: agent.Credits,
		Headquarters: agent.Headquarters,
		DB: db,
	}, nil
}
