package api

import (
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

type apiConfig struct {
	db *data.Queries
}

func NewService(dbQueries *data.Queries, port string) *http.Server {
	cfg := apiConfig{db: dbQueries}
	mux := http.NewServeMux()

	//Handlers
	mux.HandleFunc("GET /api/health", cfg.handlerHealth)

	// Games
	mux.HandleFunc("GET /api/games", cfg.handlerGetGames)
	mux.HandleFunc("POST /api/games", cfg.handlerAddGame)
	mux.HandleFunc("DELETE /api/games/{gameID}", cfg.handlerDeleteGame)

	server := http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	return &server
}
