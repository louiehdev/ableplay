package api

import (
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (api *apiConfig) handlerGetGamesChanges(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	gamesChanges, err := api.DB.GetGamesChanges(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve potential games from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, gamesChanges)
}

func (api *apiConfig) handlerGetFeaturesChanges(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	featuresChanges, err := api.DB.GetFeaturesChanges(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve potential features from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, featuresChanges)
}

