package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/louiehdev/ableplay/internal/data"
)

func (f *frontendConfig) handlerAddGameForm(w http.ResponseWriter, _ *http.Request) {
	f.Templates.ExecuteTemplate(w, "addGameForm", nil)
}

func (f *frontendConfig) handlerAddGamePlatform(w http.ResponseWriter, _ *http.Request) {
	f.Templates.ExecuteTemplate(w, "addGamePlatform", nil)
}

func (f *frontendConfig) handlerGameList(w http.ResponseWriter, _ *http.Request) {
	resp, err := http.Get(f.APIBase + "/api/games")
	if err != nil {
		http.Error(w, "failed to fetch games", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var games []data.GetGamesRow
	json.NewDecoder(resp.Body).Decode(&games)
	var gamesList []data.GameParamsJSON
	for _, game := range games {
		gamesList = append(gamesList, data.GameParamsJSON{
			Title:       game.Title,
			Developer:   game.Developer.String,
			Publisher:   game.Publisher.String,
			ReleaseYear: strconv.Itoa(int(game.ReleaseYear.Int32)),
			Platforms:   game.Platforms,
			Description: game.Description.String})
	}
	f.Templates.ExecuteTemplate(w, "gameList", gamesList)
}
