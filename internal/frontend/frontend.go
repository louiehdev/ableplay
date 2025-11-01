package frontend

import (
	"bytes"
	"context"
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

type frontendConfig struct {
	templates  *template.Template
	httpClient *http.Client
	apiBase    string
}

func NewService(tmpl *template.Template, apibase string) *http.ServeMux {
	cfg := frontendConfig{templates: tmpl, httpClient: &http.Client{}, apiBase: apibase}
	mux := http.NewServeMux()

	// Handlers
	mux.HandleFunc("GET /", cfg.handlerHome)
	mux.HandleFunc("GET /initialize", cfg.handlerInitialize)

	// Games
	mux.HandleFunc("GET /games", cfg.handlerFrontendGames)
	mux.HandleFunc("GET /games/addformplatform", cfg.handlerAddGamePlatform)
	mux.HandleFunc("GET /games/newform", cfg.handlerAddGameForm)
	mux.HandleFunc("GET /games/updateform", cfg.handlerUpdateGameForm)
	mux.HandleFunc("GET /games/list", cfg.handlerFrontendGetGames)
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

	return mux
}

func (f *frontendConfig) callAPI(ctx context.Context, method, path string, payload any) (*http.Response, error) {
	var body io.Reader
	if payload != nil {
		b, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(b)
	}
	req, err := http.NewRequestWithContext(ctx, method, f.apiBase+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
