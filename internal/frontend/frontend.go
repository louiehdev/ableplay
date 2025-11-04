package frontend

import (
	"embed"
	"html/template"
	"net/http"
)

//go:embed templates/*.html
var templates embed.FS

type frontendConfig struct {
	templates  *template.Template
	httpClient *http.Client
	apiBase    string
	platform   string
}

func NewService(tmpl *template.Template, apibase, platform string) *http.ServeMux {
	cfg := frontendConfig{templates: tmpl, httpClient: &http.Client{}, apiBase: apibase, platform: platform}
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("GET /", cfg.handlerHome)
	mux.HandleFunc("GET /demo", cfg.handlerInitializeDemoData)

	// Games
	mux.HandleFunc("GET /games/", cfg.handlerFrontendGameFeatures)
	mux.HandleFunc("GET /games/addformplatform", cfg.handlerAddGamePlatform)
	mux.HandleFunc("GET /games/newform", cfg.handlerAddGameForm)
	mux.HandleFunc("GET /games/updateform", cfg.handlerUpdateGameForm)
	mux.HandleFunc("POST /games/add", cfg.handlerFrontendAddGame)
	mux.HandleFunc("PUT /games/update", cfg.handlerFrontendUpdateGame)
	mux.HandleFunc("DELETE /games/delete", cfg.handlerFrontendDeleteGame)

	// Features
	mux.HandleFunc("GET /features", cfg.handlerFrontendFeatures)
	mux.HandleFunc("GET /features/newform", cfg.handlerAddFeatureForm)
	mux.HandleFunc("GET /features/updateform", cfg.handlerUpdateFeatureForm)
	mux.HandleFunc("GET /features/list", cfg.handlerFrontendGetFeatures)
	mux.HandleFunc("POST /features/add", cfg.handlerFrontendAddFeature)
	mux.HandleFunc("PUT /features/update", cfg.handlerFrontendUpdateFeature)
	mux.HandleFunc("DELETE /features/delete", cfg.handlerFrontendDeleteFeature)

	// Game Features
	mux.HandleFunc("GET /games/feature", cfg.handlerFrontendGetGameFeature)
	mux.HandleFunc("GET /games/features/newform", cfg.handlerGameFeatureForm)
	mux.HandleFunc("GET /games/features/list", cfg.handlerFrontendGetGamesFeatures)
	mux.HandleFunc("POST /games/features/add", cfg.handlerFrontendAddGameFeature)

	return mux
}

func LoadTemplates() *template.Template {
	return template.Must(template.ParseFS(templates, "templates/*.html"))
}
