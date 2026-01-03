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
	mux.HandleFunc("POST /api/games", api.RequireRole("admin")(api.handlerAddGame))
	mux.HandleFunc("PUT /api/games/{gameID}", api.RequireRole("admin")(api.handlerUpdateGame))
	mux.HandleFunc("DELETE /api/games/{gameID}", api.RequireRole("admin")(api.handlerDeleteGame))

	// Features
	mux.HandleFunc("GET /api/features/{featureID}", api.handlerGetFeature)
	mux.HandleFunc("GET /api/features", api.handlerGetFeatures)
	mux.HandleFunc("GET /api/features/search", api.handlerSearchFeatures)
	mux.HandleFunc("POST /api/features", api.RequireRole("admin")(api.handlerAddFeature))
	mux.HandleFunc("PUT /api/features/{featureID}", api.RequireRole("admin")(api.handlerUpdateFeature))
	mux.HandleFunc("DELETE /api/features/{featureID}", api.RequireRole("admin")(api.handlerDeleteFeature))

	// Game Features
	mux.HandleFunc("GET /api/games/features", api.handlerGetGamesWithFeatures)
	mux.HandleFunc("GET /api/games/{gameID}/features", api.handlerGetFeaturesByGame)
	mux.HandleFunc("GET /api/features/{featureID}/games", api.handlerGetGamesByFeature)
	mux.HandleFunc("GET /api/games/{gameID}/features/{featureID}", api.handlerGetGameFeature)
	mux.HandleFunc("POST /api/games/{gameID}/features", api.RequireRole("admin")(api.handlerAddGameFeature))
	mux.HandleFunc("PUT /api/games/{gameID}/features/{featureID}", api.RequireRole("admin")(api.handlerUpdateGameFeature))
	mux.HandleFunc("DELETE /api/games/{gameID}/features/{featureID}", api.RequireRole("admin")(api.handlerDeleteGameFeature))

	// Changes
	mux.HandleFunc("GET /api/changes/games", api.RequireRole("moderator")(api.handlerGetGamesChanges))
	mux.HandleFunc("GET /api/changes/features", api.RequireRole("moderator")(api.handlerGetFeaturesChanges))
	mux.HandleFunc("GET /api/changes/gamesfeatures", api.RequireRole("moderator")(api.handlerGetGamesFeaturesChanges))
	mux.HandleFunc("POST /api/changes/games", api.RequireRole("moderator")(api.handlerAddGameChange))
	mux.HandleFunc("POST /api/changes/features", api.RequireRole("moderator")(api.handlerAddFeatureChange))
	mux.HandleFunc("POST /api/changes/gamesfeatures", api.RequireRole("moderator")(api.handlerAddGameFeatureChange))
	mux.HandleFunc("PUT /api/changes/games", api.RequireRole("moderator")(api.handlerSubmitChanges))
	mux.HandleFunc("PUT /api/changes/features", api.RequireRole("moderator")(api.handlerSubmitChanges))
	mux.HandleFunc("PUT /api/changes/gamesfeatures", api.RequireRole("moderator")(api.handlerSubmitChanges))

	// Users
	mux.HandleFunc("GET /api/users", api.RequireRole("user")(api.handlerCreateAPIKey))
	mux.HandleFunc("POST /api/users", api.handlerCreateUser)
	mux.HandleFunc("POST /api/login", api.RequireRole("user")(api.handlerLogin))
	mux.HandleFunc("PUT /api/users", api.RequireRole("user")(api.handlerUpdateUser))

	// Utils
	mux.HandleFunc("GET /api/health", api.handlerHealth)

	return mux
}
