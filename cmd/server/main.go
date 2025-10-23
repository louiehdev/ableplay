package main

import (
	"context"
	"log"

	"github.com/louiehdev/ableplay/internal/db"
)

func main() {
	ctx := context.Background()
	dbConn, err := db.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()
	if err := db.Migrate(dbConn); err != nil {
		log.Fatal(err)
	}

	if err := dbConn.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to server!")
}
