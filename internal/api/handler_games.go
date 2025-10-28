package api

import (
	"encoding/json"
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (cfg *apiConfig) handlerAddGame(w http.ResponseWriter, r *http.Request) {
	var params data.AddGameParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	_, err := cfg.DB.AddGame(r.Context(), params)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to add game to database")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (cfg *apiConfig) handlerGetGames(w http.ResponseWriter, r *http.Request) {
	games, err := cfg.DB.GetGames(r.Context())
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to receive games from database")
	}

	data.RespondWithJSON(w, http.StatusOK, games)
}

func (cfg *apiConfig) handlerDeleteGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := data.GetRequestUUID(r, "gameID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game not found")
		return
	}

	if err := cfg.DB.DeleteGame(r.Context(), gameID); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
