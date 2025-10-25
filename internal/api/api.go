package api

import (
	"net/http"

	"github.com/louiehdev/ableplay/internal/config"
)

type apiConfig struct {
	*config.AppConfig
}

func NewService(app *config.AppConfig) *http.ServeMux {
	cfg := apiConfig{AppConfig: app}
	mux := http.NewServeMux()

	// Games
	mux.HandleFunc("GET /api/games", cfg.handlerGetGames)
	mux.HandleFunc("POST /api/games", cfg.handlerAddGame)
	mux.HandleFunc("DELETE /api/games/{gameID}", cfg.handlerDeleteGame)

	// Utils
	mux.HandleFunc("GET /api/health", cfg.handlerHealth)

	return mux
}
