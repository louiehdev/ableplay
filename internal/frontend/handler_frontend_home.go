package frontend

import (
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	f.templates.ExecuteTemplate(w, "home.html", nil)
}

func (f *frontendConfig) handlerInitialize(w http.ResponseWriter, r *http.Request) {
	// Todo: Add hardcoded games to put in database

	// Features
	for _, feature := range data.GetHardcodedFeatures() {
		_, resperror := f.callAPI(r.Context(), "POST", "/api/features", feature)
		if resperror != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Failed to add feature to database")
		}
	}
	
	w.WriteHeader(http.StatusOK)
}