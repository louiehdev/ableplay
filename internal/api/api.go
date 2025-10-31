package api

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louiehdev/ableplay/internal/data"
)

type apiConfig struct {
	DB *data.Queries
}

func NewService(dbConn *pgxpool.Pool) *http.ServeMux {
	dbQueries := data.New(dbConn)
	api := apiConfig{DB: dbQueries}
	mux := http.NewServeMux()

	// Games
	mux.HandleFunc("GET /api/games/{gameID}", api.handlerGetGame)
	mux.HandleFunc("GET /api/games", api.handlerGetGames)
	mux.HandleFunc("POST /api/games", api.handlerAddGame)
	mux.HandleFunc("PUT /api/games/{gameID}", api.handlerUpdateGame)
	mux.HandleFunc("DELETE /api/games/{gameID}", api.handlerDeleteGame)

	// Features
	mux.HandleFunc("GET /api/features/{featureID}", api.handlerGetFeature)
	mux.HandleFunc("GET /api/features", api.handlerGetFeatures)
	mux.HandleFunc("POST /api/features", api.handlerAddFeature)
	mux.HandleFunc("PUT /api/features/{featureID}", api.handlerUpdateFeature)
	mux.HandleFunc("DELETE /api/features/{featureID}", api.handlerDeleteFeature)

	// Utils
	mux.HandleFunc("GET /api/health", api.handlerHealth)

	return mux
}
