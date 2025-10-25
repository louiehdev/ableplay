package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/louiehdev/ableplay/internal/api"
	"github.com/louiehdev/ableplay/internal/config"
	"github.com/louiehdev/ableplay/internal/db"
	"github.com/louiehdev/ableplay/internal/web"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	apiPort := os.Getenv("API_PORT")
	frontendPort := os.Getenv("FRONTEND_PORT")
	if dbURL == "" || apiPort == "" || frontendPort == "" {
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

	appCfg := config.NewAppConfig(dbConn, apiPort, frontendPort)

	rootMux := http.NewServeMux()
	rootMux.Handle("/api/", api.NewService(appCfg))
	rootMux.Handle("/", web.NewService(appCfg))

	server := http.Server{
		Addr:    ":" + appCfg.APIPort,
		Handler: rootMux,
	}

	log.Fatal(server.ListenAndServe())
}
