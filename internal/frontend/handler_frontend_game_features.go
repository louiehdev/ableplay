package frontend

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerFrontendGameFeatures(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "gamesPage", nil)
}

func (f *frontendConfig) handlerGameFeatureForm(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")
	gameTitle := r.URL.Query().Get("title")

	resp, resperror := f.callAPI(r.Context(), r.Method, "/api/features", nil)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch features")
		return
	}
	defer resp.Body.Close()

	var features []data.FeatureData
	json.NewDecoder(resp.Body).Decode(&features)
	var featureList []data.FeaturePublic
	for _, feature := range features {
		featureList = append(featureList, data.FeaturePublic{
			ID:   feature.ID.String(),
			Name: feature.Name})
	}

	formData := struct {
		GameID      string
		GameTitle   string
		FeatureList []data.FeaturePublic
	}{GameID: gameID, GameTitle: gameTitle, FeatureList: featureList}

	f.templates.ExecuteTemplate(w, "gamefeatureForm", formData)
}

func (f *frontendConfig) handlerFrontendAddGameFeature(w http.ResponseWriter, r *http.Request) {
	var params data.GameFeaturePublic
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	gameUUID, _ := uuid.Parse(params.GameID)
	featureUUID, _ := uuid.Parse(params.FeatureID)
	gamefeatureParams := data.CreateGameFeatureParams{
		GameID:    gameUUID,
		FeatureID: featureUUID,
		Notes:     data.ToPgtypeText(params.Notes),
		Verified:  data.IsChecked(params.Verified),
	}

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/games/"+params.GameID+"/features", gamefeatureParams)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to add game feature")
		return
	}

	w.Header().Set("HX-Trigger", "gamefeatureAdded")
	w.WriteHeader(http.StatusCreated)
}

func (f *frontendConfig) handlerFrontendGetGameFeature(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")
	featureID := r.URL.Query().Get("feature_id")

	resp, err := f.callAPI(r.Context(), r.Method, "/api/games/"+gameID+"/features/"+featureID, nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch game feature")
		return
	}
	defer resp.Body.Close()

	var gamefeatureData data.GetGameFeatureRow
	json.NewDecoder(resp.Body).Decode(&gamefeatureData)
	gameFeature := struct {
		Notes    string `json:"notes"`
		Verified bool   `json:"verified"`
		Title    string `json:"title"`
		Name     string `json:"name"`
	}{
		Notes:    gamefeatureData.Notes.String,
		Verified: gamefeatureData.Verified,
		Title:    gamefeatureData.Title,
		Name:     gamefeatureData.Name}

	f.templates.ExecuteTemplate(w, "gamefeaturePopover", gameFeature)
}

func (f *frontendConfig) handlerFrontendGetGamesFeatures(w http.ResponseWriter, r *http.Request) {
	resp, err := f.callAPI(r.Context(), r.Method, "/api/games/features", nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch games")
		return
	}
	defer resp.Body.Close()

	var gamesData []data.GetGamesWithFeaturesRow
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
	f.templates.ExecuteTemplate(w, "gamesfeaturesList", gamesList)
}

func (f *frontendConfig) handlerFrontendDeleteGameFeature(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")
	featureID := r.URL.Query().Get("feature_id")

	_, err := f.callAPI(r.Context(), r.Method, "/api/games/"+gameID+"/features/"+featureID, nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch game feature")
		return
	}

	w.Header().Set("HX-Trigger", "gamefeatureDeleted")
	w.WriteHeader(http.StatusNoContent)
}
