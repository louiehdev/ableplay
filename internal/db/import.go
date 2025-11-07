package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/louiehdev/ableplay/internal/data"
)

type RAWGGamesList struct {
	Results []struct {
		ID int `json:"id"`
	} `json:"results"`
}

type RAWGGameData struct {
	Title       string `json:"name"`
	Slug        string `json:"slug"`
	Released    string `json:"released"`
	Description string `json:"description"`
	Image       string `json:"background_image"`
	Platforms   []struct {
		Platform struct {
			Name string `json:"name"`
		} `json:"platform"`
	} `json:"platforms"`
}

func Import(dbConn *pgxpool.Pool, ctx context.Context) error {
	godotenv.Load()
	apiKey := os.Getenv("RAWG_API_KEY")
	if apiKey == "" {
		log.Fatal("Environment variables must be set")
	}
	dbQueries := data.New(dbConn)

	games, err := fetchGames(apiKey)
	if err != nil {
		return err
	}

	for _, game := range games {
		_, err := dbQueries.AddGame(ctx, game)
		if err != nil {
			log.Printf("Unable to add game '%s' into database", game.Title)
		}
	}

	for _, feature := range data.GetHardcodedFeatures() {
		_, err := dbQueries.AddFeature(ctx, feature)
		if err != nil {
			log.Printf("Unable to add feature '%s' into database", feature.Name)
		}
	}

	log.Print("Import complete")
	return nil
}

func fetchGames(apiKey string) ([]data.AddGameParams, error) {
	var rawgGamesData []data.AddGameParams
	rawgGameIDList, err := fetchGameList(apiKey)
	if err != nil {
		return rawgGamesData, err
	}

	for _, gameID := range rawgGameIDList {
		gameData, err := createGameData(apiKey, gameID)
		if err != nil {
			return rawgGamesData, err
		}
		rawgGamesData = append(rawgGamesData, gameData)
		time.Sleep(250 * time.Millisecond)
	}

	return rawgGamesData, nil
}

func fetchGameList(apiKey string) ([]int, error) {
	url := fmt.Sprintf("https://api.rawg.io/api/games?key=%s&page_size=50", apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data RAWGGamesList
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	var ids []int
	for _, game := range data.Results {
		ids = append(ids, game.ID)
	}
	return ids, nil
}

func createGameData(apiKey string, id int) (data.AddGameParams, error) {
	url := fmt.Sprintf("https://api.rawg.io/api/games/%d?key=%s", id, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return data.AddGameParams{}, nil
	}
	defer resp.Body.Close()

	var gameData RAWGGameData
	if err := json.NewDecoder(resp.Body).Decode(&gameData); err != nil {
		return data.AddGameParams{}, nil
	}
	var gamePlatforms []string
	for _, platform := range gameData.Platforms {
		if platform.Platform.Name != "" {
			gamePlatforms = append(gamePlatforms, platform.Platform.Name)
		}
	}

	gameParams := data.AddGameParams{
		Title:       gameData.Title,
		Slug:        data.ToPgtypeText(gameData.Slug),
		ReleaseYear: data.ToPgtypeInt4(strings.Split(gameData.Released, "-")[0]),
		Description: data.ToPgtypeText(parseDescription(gameData.Description)),
		Platforms:   gamePlatforms}

	return gameParams, nil
}

func parseDescription(desc string) string {
	descParagraph := strings.TrimPrefix(strings.Split(desc, "</p>")[0], "<p>")
	if len(descParagraph) == 0 {
		return ""
	}
	descShort := strings.Split(descParagraph, "<br />")[0]
	return descShort
}
