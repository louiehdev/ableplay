package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/louiehdev/ableplay/internal/api"
	"github.com/louiehdev/ableplay/internal/data"
	"github.com/louiehdev/ableplay/internal/db"
)

func main() {
	godotenv.Load()
	ctx := context.Background()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set")
	}
	dbConn, err := db.Connect(ctx, dbURL)
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
	log.Print("Successfully connected to database!")
	dbQueries := data.New(dbConn)

	apiService := api.NewService(dbQueries, port)

	log.Fatal(apiService.ListenAndServe())
}
