package frontend

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerFrontendGamesFeatures(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "gamesPage", nil)
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

	var gameFeature data.GameFeaturePublic
	json.NewDecoder(resp.Body).Decode(&gameFeature)

	f.templates.ExecuteTemplate(w, "gamefeatureCard", gameFeature)
}

func (f *frontendConfig) handlerFrontendGetGamesFeatures(w http.ResponseWriter, r *http.Request) {
	resp, err := f.callAPI(r.Context(), r.Method, "/api/games/features", nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch games")
		return
	}
	defer resp.Body.Close()

	var games []data.GamePublic
	json.NewDecoder(resp.Body).Decode(&games)

	f.templates.ExecuteTemplate(w, "gamesList", games)
}

func (f *frontendConfig) handlerFrontendGetGamesByFeature(w http.ResponseWriter, r *http.Request) {
	featureID := r.URL.Query().Get("feature_id")

	resp, err := f.callAPI(r.Context(), r.Method, "/api/features/"+featureID+"/games", nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch games using feature id")
		return
	}
	defer resp.Body.Close()

	var games []data.GameByFeaturePublic
	json.NewDecoder(resp.Body).Decode(&games)

	if err := f.templates.ExecuteTemplate(w, "gamesbyfeatureList", games); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
}

func (f *frontendConfig) handlerAddGameFeatureForm(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("game_id")
	gameTitle := r.URL.Query().Get("title")

	resp, err := f.callAPI(r.Context(), r.Method, "/api/features", nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch features")
		return
	}
	defer resp.Body.Close()

	var features []data.FeaturePublic
	json.NewDecoder(resp.Body).Decode(&features)

	formData := struct {
		GameID      string
		GameTitle   string
		FeatureList []data.FeaturePublic
	}{GameID: gameID, GameTitle: gameTitle, FeatureList: features}

	f.templates.ExecuteTemplate(w, "addGameFeatureForm", formData)
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
		Notes:     params.Notes,
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

func (f *frontendConfig) handlerFrontendDeleteGameFeature(w http.ResponseWriter, r *http.Request) {
	gameID := r.PathValue("gameID")
	featureID := r.PathValue("featureID")

	_, err := f.callAPI(r.Context(), "DELETE", "/api/games/"+gameID+"/features/"+featureID, nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to delete game feature")
		return
	}

	w.Header().Set("HX-Trigger", "gamefeatureDeleted")
	w.WriteHeader(http.StatusOK)
}
