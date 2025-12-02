package api

import (
	"encoding/json"
	"net/http"

	"github.com/louiehdev/ableplay/internal/data"
)

func (api *apiConfig) handlerAddGame(w http.ResponseWriter, r *http.Request) {
	var params data.AddGameParams
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	_, err := api.DB.AddGame(r.Context(), params)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to add game to database")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (api *apiConfig) handlerGetGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := data.GetRequestUUID(r, "gameID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game not found")
		return
	}

	game, err := api.DB.GetGame(r.Context(), gameID)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve game")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, game)
}

func (api *apiConfig) handlerGetGames(w http.ResponseWriter, r *http.Request) {
	queryParams := data.ParseQueryParams(r.URL.Query())
	limit, _ := queryParams["limit"].(int32)

	games, err := api.DB.GetGames(r.Context(), limit)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve games from database")
		return
	}

	data.RespondWithJSON(w, http.StatusOK, games)
}

func (api *apiConfig) handlerSearchGames(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")

	games, err := api.DB.GetGamesSearch(r.Context(), query)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to find matching games from database")
		return
	}

	var response []data.GameData
	for _, game := range games {
		var gameFeatures []data.GameFeatureData
		json.Unmarshal(game.GameFeatures, &gameFeatures)
		response = append(response, data.GameData{
			ID:           game.ID,
			Title:        game.Title,
			Developer:    game.Developer,
			Publisher:    game.Publisher,
			ReleaseYear:  game.ReleaseYear,
			Platforms:    game.Platforms,
			Description:  game.Description,
			GameFeatures: gameFeatures})
	}

	data.RespondWithJSON(w, http.StatusOK, response)
}

func (api *apiConfig) handlerUpdateGame(w http.ResponseWriter, r *http.Request) {
	var params data.UpdateGameParams
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	if err := api.DB.UpdateGame(r.Context(), params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Unable to update feature")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (api *apiConfig) handlerDeleteGame(w http.ResponseWriter, r *http.Request) {
	gameID, err := data.GetRequestUUID(r, "gameID")
	if err != nil {
		data.RespondWithError(w, http.StatusNotFound, "Game not found")
		return
	}

	if err := api.DB.DeleteGame(r.Context(), gameID); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
