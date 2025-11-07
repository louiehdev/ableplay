package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (api *apiConfig) handlerAddGameFeature(w http.ResponseWriter, r *http.Request) {
	var params data.CreateGameFeatureParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	_, err := api.DB.CreateGameFeature(r.Context(), params)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to add game feature to database")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *apiConfig) handlerDeleteGameFeature(w http.ResponseWriter, r *http.Request) {
	gameID, err := data.GetRequestUUID(r, "gameID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game feature not found")
		return
	}
	featureID, err := data.GetRequestUUID(r, "featureID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game feature not found")
		return
	}

	if err := api.DB.DeleteGameFeature(r.Context(), data.DeleteGameFeatureParams{GameID: gameID, FeatureID: featureID}); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to delete game feature from database")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *apiConfig) handlerGetGamesWithFeatures(w http.ResponseWriter, r *http.Request) {
	games, err := api.DB.GetGamesWithFeatures(r.Context())
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Unable to receive games and features from database, %v", err))
		return
	}

	data.RespondWithJSON(w, http.StatusOK, games)
}

func (api *apiConfig) handlerGetGameFeature(w http.ResponseWriter, r *http.Request) {
	gameID, err := data.GetRequestUUID(r, "gameID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game feature not found")
		return
	}
	featureID, err := data.GetRequestUUID(r, "featureID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game feature not found")
		return
	}

	gameFeature, err := api.DB.GetGameFeature(r.Context(), data.GetGameFeatureParams{GameID: gameID, FeatureID: featureID})
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve game feature from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, gameFeature)
}

func (api *apiConfig) handlerGetFeaturesByGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := data.GetRequestUUID(r, "gameID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game not found")
		return
	}

	gameFeatures, err := api.DB.GetFeaturesByGame(r.Context(), gameID)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve features by game from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, gameFeatures)
}

func (api *apiConfig) handlerGetGamesByFeature(w http.ResponseWriter, r *http.Request) {
	featureID, err := data.GetRequestUUID(r, "featureID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Feature not found")
		return
	}

	games, err := api.DB.GetGamesByFeature(r.Context(), featureID)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to receive games by feature from database")
	}

	data.RespondWithJSON(w, http.StatusOK, games)
}

func (api *apiConfig) handlerUpdateGameFeature(w http.ResponseWriter, r *http.Request) {
	var params data.UpdateGameFeatureParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if err := api.DB.UpdateGameFeature(r.Context(), params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to update feature")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
