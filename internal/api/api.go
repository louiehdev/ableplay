package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louiehdev/ableplay/internal/data"
)

type apiConfig struct {
	DB *data.Queries
	Conn *pgxpool.Pool
}

func NewService(dbConn *pgxpool.Pool) *http.ServeMux {
	dbQueries := data.New(dbConn)
	api := apiConfig{DB: dbQueries, Conn: dbConn}
	mux := http.NewServeMux()

	// Games
	mux.HandleFunc("GET /api/games/{gameID}", api.handlerGetGame)
	mux.HandleFunc("GET /api/games", api.handlerGetGames)
	mux.HandleFunc("GET /api/games/search", api.handlerSearchGames)
	mux.HandleFunc("POST /api/games", api.handlerAddGame)
	mux.HandleFunc("PUT /api/games/{gameID}", api.handlerUpdateGame)
	mux.HandleFunc("DELETE /api/games/{gameID}", api.handlerDeleteGame)

	// Features
	mux.HandleFunc("GET /api/features/{featureID}", api.handlerGetFeature)
	mux.HandleFunc("GET /api/features", api.handlerGetFeatures)
	mux.HandleFunc("GET /api/features/search", api.handlerSearchFeatures)
	mux.HandleFunc("POST /api/features", api.handlerAddFeature)
	mux.HandleFunc("PUT /api/features/{featureID}", api.handlerUpdateFeature)
	mux.HandleFunc("DELETE /api/features/{featureID}", api.handlerDeleteFeature)

	// Game Features
	mux.HandleFunc("GET /api/games/features", api.handlerGetGamesWithFeatures)
	mux.HandleFunc("GET /api/games/{gameID}/features", api.handlerGetFeaturesByGame)
	mux.HandleFunc("GET /api/features/{featureID}/games", api.handlerGetGamesByFeature)
	mux.HandleFunc("GET /api/games/{gameID}/features/{featureID}", api.handlerGetGameFeature)
	mux.HandleFunc("POST /api/games/{gameID}/features", api.handlerAddGameFeature)
	mux.HandleFunc("PUT /api/games/{gameID}/features/{featureID}", api.handlerUpdateGameFeature)
	mux.HandleFunc("DELETE /api/games/{gameID}/features/{featureID}", api.handlerDeleteGameFeature)

	// Utils
	mux.HandleFunc("GET /api/health", api.handlerHealth)

	return mux
}
