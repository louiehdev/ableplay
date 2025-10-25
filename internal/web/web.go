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

	return mux
}
