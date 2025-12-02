package api

import (
	"encoding/json"
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (api *apiConfig) handlerAddFeature(w http.ResponseWriter, r *http.Request) {
	var params data.AddFeatureParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	_, err := api.DB.AddFeature(r.Context(), params)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to add feature to database")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *apiConfig) handlerGetFeature(w http.ResponseWriter, r *http.Request) {
	featureID, err := data.GetRequestUUID(r, "featureID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Feature not found")
		return
	}

	feature, err := api.DB.GetFeature(r.Context(), featureID)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, feature)
}

func (api *apiConfig) handlerGetFeatures(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	features, err := api.DB.GetFeatures(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve features from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, features)
}

func (api *apiConfig) handlerSearchFeatures(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	features, err := api.DB.GetFeaturesSearch(r.Context(), query)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to find matching features from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, features)
}

func (api *apiConfig) handlerUpdateFeature(w http.ResponseWriter, r *http.Request) {
	var params data.UpdateFeatureParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if err := api.DB.UpdateFeature(r.Context(), params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to update feature")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *apiConfig) handlerDeleteFeature(w http.ResponseWriter, r *http.Request) {
	featureID, err := data.GetRequestUUID(r, "featureID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Feature not found")
		return
	}

	if err := api.DB.DeleteFeature(r.Context(), featureID); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
