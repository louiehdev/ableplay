package api

import (
	"encoding/json"
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (cfg *apiConfig) handlerAddGame(w http.ResponseWriter, r *http.Request) {
	var params data.AddGameParamsJSON
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	gameParams := data.AddGameParams{
		Title:       params.Title,
		Developer:   toPgtypeText(params.Developer),
		Publisher:   toPgtypeText(params.Publisher),
		ReleaseYear: toPgtypeInt4(params.ReleaseYear),
		Platforms:   params.Platforms,
		Description: toPgtypeText(params.Description),
	}

	game, err := cfg.db.AddGame(r.Context(), gameParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to add game to database")
		return
	}

	respondWithJSON(w, http.StatusCreated, game)
}

func (cfg *apiConfig) handlerGetGames(w http.ResponseWriter, r *http.Request) {
	games, err := cfg.db.GetGames(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Unable to receive games from database")
	}

	respondWithJSON(w, http.StatusOK, games)
}

func (cfg *apiConfig) handlerDeleteGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := getRequestUUID(r, "gameID")
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Game not found")
		return
	}

	if err := cfg.db.DeleteGame(r.Context(), gameID); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respondWithError(w, http.StatusNoContent, "Game deleted")
}
