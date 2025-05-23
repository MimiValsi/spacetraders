package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/MimiValsi/spacetraders/internal/database"
	"github.com/MimiValsi/spacetraders/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	accountToken := os.Getenv("ACCOUNT")
	if accountToken == "" {
		log.Fatalln("Must set ACCOUNT token")
	}

	dbURL := os.Getenv("dbURL")
	if dbURL == "" {
		log.Fatalln("Must set dbURL token")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("couldn't open DB: %s", err)
	}

	queries := database.New(db)

	client, err := api.NewClient(context.Background(), accountToken, queries)
	if err != nil {
		log.Fatalln("new client error")
	}

	if err = client.Register("MIMI8", "AEGIS"); err != nil {
		log.Fatalln(err)
	}

}
