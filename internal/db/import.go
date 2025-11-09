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

	for _, feature := range getHardcodedFeatures() {
		_, err := dbQueries.AddFeature(ctx, feature)
		if err != nil {
			log.Printf("Unable to add feature '%s' into database with error: %v", feature.Name, err)
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

func getHardcodedFeatures() []data.AddFeatureParams {
	return []data.AddFeatureParams{
		CreateHardcodedFeature("Subtitles", "Display text for dialogue and important sound effects", "Visual"),
		CreateHardcodedFeature("Colorblind Modes", "Provide filters or distinct color palettes to help players with color vision deficiencies", "Visual"),
		CreateHardcodedFeature("Text Formatting", "Allow players to adjust size or style of text on screen", "Visual"),
		CreateHardcodedFeature("Visual Cues", "Provide visual indicators for in-game sounds", "Visual"),
		CreateHardcodedFeature("High Contrast", "Enhance visibility by increasing the contrast between foreground and background elements", "Visual"),
		CreateHardcodedFeature("Screen Magnifier", "Allow players to zoom in or out on screen elements", "Visual"),
		CreateHardcodedFeature("Volume Controls", "Allow players to adjust the volume for independent sound elements such as dialogue, sound effects, or music", "Audio"),
		CreateHardcodedFeature("Audio Cues", "Provide sound effects for important gameplay events", "Audio"),
		CreateHardcodedFeature("Text to Speech", "Read on-screen text aloud", "Audio"),
		CreateHardcodedFeature("Remappable Controls", "Allow players to change game inputs such as buttons or keys, or to toggle actions on or off instead of holding a button down", "Motor"),
		CreateHardcodedFeature("Adjustable Difficulty", "Offer a range of difficulty settings, including assist modes that might provide more resources or slow down gameplay", "Motor"),
		CreateHardcodedFeature("Alternative Input", "Support for additional game controllers or input schemes", "Motor"),
		CreateHardcodedFeature("Game Speed", "Provide an option to change the speed at which the game runs", "Motor"),
		CreateHardcodedFeature("Clear Interfaces", "Menus and in-game information is clear and easy to navigate", "Cognitive"),
		CreateHardcodedFeature("Motion Blur and Screen Shake", "Options to disable or reduce visual effects that can cause motion sickness", "Cognitive"),
	}
}

func CreateHardcodedFeature(name, description, category string) data.AddFeatureParams {
	return data.AddFeatureParams{Name: name, Description: data.ToPgtypeText(description), Category: category}
}
