package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/louiehdev/ableplay/internal/frontend"
)

func main() {
	godotenv.Load()
	port := os.Getenv("WEB_PORT")
	apiBase := os.Getenv("API_BASE")
	platform := os.Getenv("PLATFORM")
	if port == "" || apiBase == "" || platform == "" {
		log.Fatal("Environment variables must be set")
	}

	tmpl := template.Must(template.ParseGlob("./internal/frontend/templates/*"))

	server := http.Server{
		Addr:    ":" + port,
		Handler: frontend.NewService(tmpl, apiBase, platform),
	}

	log.Print("Frontend running successfully")
	log.Fatal(server.ListenAndServe())
}
