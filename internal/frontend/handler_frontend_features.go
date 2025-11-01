package frontend

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerFrontendFeatures(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "featuresPage", nil)
}

func (f *frontendConfig) handlerAddFeatureForm(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "addFeatureForm", nil)
}

func (f *frontendConfig) handlerUpdateFeatureForm(w http.ResponseWriter, r *http.Request) {
	featureID := r.URL.Query().Get("id")

	resp, resperror := f.callAPI(r.Context(), r.Method, "/api/features/"+featureID, nil)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch feature")
		return
	}
	defer resp.Body.Close()

	var featureData data.FeatureData
	if err := json.NewDecoder(resp.Body).Decode(&featureData); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong: decoding failed")
		return
	}

	feature := data.FeaturePublic{
		ID:          featureData.ID.String(),
		Name:        featureData.Name,
		Description: featureData.Description.String,
		Category:    featureData.Category.String}

	f.templates.ExecuteTemplate(w, "updateFeatureForm", feature)
}

func (f *frontendConfig) handlerFrontendAddFeature(w http.ResponseWriter, r *http.Request) {
	var params data.FeaturePublic
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	addFeatureParams := data.AddFeatureParams{
		Name:        params.Name,
		Description: data.ToPgtypeText(params.Description),
		Category:    data.ToPgtypeText(params.Category)}

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/features", addFeatureParams)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to add feature")
		return
	}

	w.Header().Set("HX-Trigger", "featureAdded")
	w.WriteHeader(http.StatusCreated)
}

func (f *frontendConfig) handlerFrontendUpdateFeature(w http.ResponseWriter, r *http.Request) {
	var params data.FeaturePublic
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	featureUUID, _ := uuid.Parse(params.ID)
	updateFeatureParams := data.FeatureData{
		ID:          featureUUID,
		Name:        params.Name,
		Description: data.ToPgtypeText(params.Description),
		Category:    data.ToPgtypeText(params.Category),
	}

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/features/"+params.ID, updateFeatureParams)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to update feature")
		return
	}

	w.Header().Set("HX-Trigger", "featureUpdated")
	w.WriteHeader(http.StatusNoContent)
}

func (f *frontendConfig) handlerFrontendGetFeatures(w http.ResponseWriter, r *http.Request) {
	resp, err := f.callAPI(r.Context(), r.Method, "/api/features", nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch features")
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

	f.templates.ExecuteTemplate(w, "featureList", featureList)
}

func (f *frontendConfig) handlerFrontendDeleteFeature(w http.ResponseWriter, r *http.Request) {
	featureID := r.URL.Query().Get("id")

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/features/"+featureID, nil)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to delete feature")
		return
	}

	w.Header().Set("HX-Trigger", "featureDeleted")
	w.WriteHeader(http.StatusNoContent)
}
