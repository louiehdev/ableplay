package frontend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerAddGameForm(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "addGameForm", nil)
}

func (f *frontendConfig) handlerAddGamePlatform(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "addGamePlatform", nil)
}

func (f *frontendConfig) handlerUpdateGameForm(w http.ResponseWriter, r *http.Request) {
	gameID := r.URL.Query().Get("id")

	resp, resperror := f.callAPI(r.Context(), r.Method, "/api/games/"+gameID, nil)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch game")
		return
	}
	defer resp.Body.Close()

	var gameData data.GameData
	if err := json.NewDecoder(resp.Body).Decode(&gameData); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Something went wrong: decoding failed for game: %v", gameID))
		return
	}

	game := data.GamePublic{
		ID:          gameData.ID.String(),
		Title:       gameData.Title,
		Developer:   gameData.Developer.String,
		Publisher:   gameData.Publisher.String,
		ReleaseYear: strconv.Itoa(int(gameData.ReleaseYear.Int32)),
		Platforms:   gameData.Platforms,
		Description: gameData.Description.String}

	f.templates.ExecuteTemplate(w, "updateGameForm", game)
}

func (f *frontendConfig) handlerFrontendAddGame(w http.ResponseWriter, r *http.Request) {
	var params data.GamePublic
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	addGameParams := data.AddGameParams{
		Title:       params.Title,
		Developer:   data.ToPgtypeText(params.Developer),
		Publisher:   data.ToPgtypeText(params.Publisher),
		ReleaseYear: data.ToPgtypeInt4(params.ReleaseYear),
		Platforms:   data.RemoveEmptyValues(params.Platforms),
		Description: data.ToPgtypeText(params.Description),
	}

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/games", addGameParams)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to add game")
		return
	}

	w.Header().Set("HX-Trigger", "gameAdded")
	w.WriteHeader(http.StatusCreated)
}

func (f *frontendConfig) handlerFrontendUpdateGame(w http.ResponseWriter, r *http.Request) {
	var params data.GamePublic
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	gameUUID, _ := uuid.Parse(params.ID)
	updateGameParams := data.GameData{
		ID:          gameUUID,
		Title:       params.Title,
		Developer:   data.ToPgtypeText(params.Developer),
		Publisher:   data.ToPgtypeText(params.Publisher),
		ReleaseYear: data.ToPgtypeInt4(params.ReleaseYear),
		Platforms:   data.RemoveEmptyValues(params.Platforms),
		Description: data.ToPgtypeText(params.Description),
	}

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/games/"+params.ID, updateGameParams)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to update game")
		return
	}

	w.Header().Set("HX-Trigger", "gameUpdated")
	w.WriteHeader(http.StatusNoContent)
}

func (f *frontendConfig) handlerFrontendDeleteGame(w http.ResponseWriter, r *http.Request) {
	gameID := r.PathValue("gameID")

	_, err := f.callAPI(r.Context(), r.Method, "/api/games/"+gameID, nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to delete game")
		return
	}

	w.Header().Set("HX-Trigger", "gameDeleted")
	w.WriteHeader(http.StatusNoContent)
}
