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
	cfg := apiConfig{DB: dbQueries}
	mux := http.NewServeMux()

	// Games
	mux.HandleFunc("GET /api/games", cfg.handlerGetGames)
	mux.HandleFunc("POST /api/games", cfg.handlerAddGame)
	mux.HandleFunc("DELETE /api/games/{gameID}", cfg.handlerDeleteGame)

	// Utils
	mux.HandleFunc("GET /api/health", cfg.handlerHealth)

	return mux
}
