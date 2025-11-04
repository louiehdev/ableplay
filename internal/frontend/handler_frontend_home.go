package frontend

import (
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	data := struct{ Platform string }{Platform: f.platform}

	f.templates.ExecuteTemplate(w, "home.html", data)
}

func (f *frontendConfig) handlerInitializeDemoData(w http.ResponseWriter, r *http.Request) {
	// Todo: Add hardcoded games to put in database
	for _, game := range data.GetHardcodedGames() {
		_, resperror := f.callAPI(r.Context(), "POST", "/api/games", game)
		if resperror != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Failed to add game to database")
		}
	}

	// Features
	for _, feature := range data.GetHardcodedFeatures() {
		_, resperror := f.callAPI(r.Context(), "POST", "/api/features", feature)
		if resperror != nil {
			data.RespondWithError(w, http.StatusInternalServerError, "Failed to add feature to database")
		}
	}

	w.WriteHeader(http.StatusOK)
}
