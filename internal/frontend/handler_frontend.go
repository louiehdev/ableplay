package frontend

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	f.templates.ExecuteTemplate(w, "home.html", nil)
}

func (f *frontendConfig) handlerDocumentation(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "apiDocs", nil)
}

func (f *frontendConfig) handlerSearch(w http.ResponseWriter, r *http.Request) {
	searchType := r.URL.Query().Get("type")
	query := r.URL.Query().Get("query")

	switch searchType {
	case "games":
		resp, err := f.callAPI(r.Context(), r.Method, "/api/games/search?q="+query, nil)
		if err != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Failed to find any matching games")
			return
		}
		defer resp.Body.Close()

		var gamesData []data.GetGamesSearchRow
		json.NewDecoder(resp.Body).Decode(&gamesData)
		var gamesList []data.GamePublic
		for _, game := range gamesData {
			var gameFeatures []data.GameFeatureData
			json.Unmarshal(game.GameFeatures, &gameFeatures)
			gamesList = append(gamesList, data.GamePublic{
				ID:           game.ID.String(),
				Title:        game.Title,
				Developer:    game.Developer.String,
				Publisher:    game.Publisher.String,
				ReleaseYear:  strconv.Itoa(int(game.ReleaseYear.Int32)),
				Platforms:    game.Platforms,
				Description:  game.Description.String,
				GameFeatures: gameFeatures})
		}
		f.templates.ExecuteTemplate(w, "gamesList", gamesList)
	case "features":
		resp, err := f.callAPI(r.Context(), r.Method, "/api/features/search?q="+query, nil)
		if err != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Failed to find any matching features")
			return
		}
		defer resp.Body.Close()

		var features []data.FeatureData
		json.NewDecoder(resp.Body).Decode(&features)
		var featureList []data.FeaturePublic
		for _, feature := range features {
			featureList = append(featureList, data.FeaturePublic{
				ID:          feature.ID.String(),
				Name:        feature.Name,
				Description: feature.Description.String,
				Category:    feature.Category.String})
		}

		f.templates.ExecuteTemplate(w, "featuresList", featureList)
	}
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
