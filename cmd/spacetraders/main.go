package main

import (
	"context"
	"log"
	"os"

	"github.com/MimiValsi/spacetraders/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	accountToken := os.Getenv("ACCOUNT")
	if accountToken == "" {
		log.Fatalln("Must set ACCOUNT token")
	}

	client, err := api.NewClient(context.Background(), accountToken)
	if err != nil {
		log.Fatalln("new client error")
	}

	if err := client.Register("MIMI3", "AEGIS"); err != nil {
		log.Fatalln(err)
	}

}
