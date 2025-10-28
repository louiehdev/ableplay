package frontend

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerAddGameForm(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "addGameForm", nil)
}

func (f *frontendConfig) handlerAddGamePlatform(w http.ResponseWriter, _ *http.Request) {
	f.templates.ExecuteTemplate(w, "addGamePlatform", nil)
}

func (f *frontendConfig) handlerFrontendAddGame(w http.ResponseWriter, r *http.Request) {
	var params data.GamePublic
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	var gamePlatforms []string
	for _, platform := range params.Platforms {
		if platform != "" {
			gamePlatforms = append(gamePlatforms, platform)
		}
	}

	addGameParams := data.AddGameParams{
		Title:       params.Title,
		Developer:   data.ToPgtypeText(params.Developer),
		Publisher:   data.ToPgtypeText(params.Publisher),
		ReleaseYear: data.ToPgtypeInt4(params.ReleaseYear),
		Platforms:   gamePlatforms,
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

func (f *frontendConfig) handlerFrontendDeleteGame(w http.ResponseWriter, r *http.Request) {
	var params struct {
		ID string `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	_, resperror := f.callAPI(r.Context(), r.Method, "/api/games/"+params.ID, nil)
	if resperror != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to delete game")
		return
	}

	w.Header().Set("HX-Trigger", "gameDeleted")
	w.WriteHeader(http.StatusNoContent)
}

func (f *frontendConfig) handlerFrontendGetGames(w http.ResponseWriter, r *http.Request) {
	resp, err := f.callAPI(r.Context(), r.Method, "/api/games", nil)
	if err != nil {
		data.RespondWithError(w, http.StatusInternalServerError, "Failed to fetch games")
		return
	}
	defer resp.Body.Close()

	var games []data.GetGamesRow
	json.NewDecoder(resp.Body).Decode(&games)
	var gamesList []data.GamePublic
	for _, game := range games {
		gamesList = append(gamesList, data.GamePublic{
			ID:          game.ID.String(),
			Title:       game.Title,
			Developer:   game.Developer.String,
			Publisher:   game.Publisher.String,
			ReleaseYear: strconv.Itoa(int(game.ReleaseYear.Int32)),
			Platforms:   game.Platforms,
			Description: game.Description.String})
	}

	f.templates.ExecuteTemplate(w, "gameList", gamesList)
}
