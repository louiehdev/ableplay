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
	mux.HandleFunc("GET /games/addformfeature", cfg.handlerAddGamePlatform)
	mux.HandleFunc("GET /games/newform", cfg.handlerAddGameForm)
	mux.HandleFunc("GET /games/updateform", cfg.handlerUpdateGameForm)
	mux.HandleFunc("GET /games/list", cfg.handlerFrontendGetGames)
	mux.HandleFunc("POST /games/addgame", cfg.handlerFrontendAddGame)
	mux.HandleFunc("PUT /games/updategame", cfg.handlerFrontendUpdateGame)
	mux.HandleFunc("DELETE /games/deletegame", cfg.handlerFrontendDeleteGame)

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
