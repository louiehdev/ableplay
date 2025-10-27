package web

import (
	"net/http"

	"github.com/louiehdev/ableplay/internal/config"
)

type frontendConfig struct {
	*config.AppConfig
}

func NewService(app *config.AppConfig) *http.ServeMux {
	cfg := frontendConfig{AppConfig: app}
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("GET /", cfg.handlerHome)
	mux.HandleFunc("GET /games/updateform", cfg.handlerAddGamePlatform)
	mux.HandleFunc("GET /games/newform", cfg.handlerAddGameForm)
	mux.HandleFunc("GET /games/list", cfg.handlerGameList)

	return mux
}
