package config

import (
	"html/template"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louiehdev/ableplay/internal/data"
)

type AppConfig struct {
	DB           *data.Queries
	Templates    *template.Template
	APIPort      string
	FrontendPort string
}

func NewAppConfig(dbConn *pgxpool.Pool, apiPort, frontendPort string) *AppConfig {
	dbQueries := data.New(dbConn)
	tmpl := template.Must(template.ParseGlob("./internal/web/templates/*"))

	return &AppConfig{DB: dbQueries, Templates: tmpl, APIPort: apiPort, FrontendPort: frontendPort}
}
