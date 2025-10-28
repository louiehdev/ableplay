package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/louiehdev/ableplay/internal/api"
	"github.com/louiehdev/ableplay/internal/db"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("API_PORT")
	if dbURL == "" || port == "" {
		log.Fatal("Environment variables must be set")
	}
	ctx := context.Background()
	dbConn, err := db.Connect(ctx, dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	if err := db.Migrate(dbConn); err != nil {
		log.Fatal(err)
	}
	log.Print("Successfully connected to database!")

	server := http.Server{
		Addr:    ":" + port,
		Handler: api.NewService(dbConn),
	}

	log.Fatal(server.ListenAndServe())
}
