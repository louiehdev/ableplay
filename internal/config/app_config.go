package config

import (
	"html/template"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/louiehdev/ableplay/internal/data"
)

type AppConfig struct {
	DB        *data.Queries
	Templates *template.Template
	Port      string
	APIBase   string
}

func NewAppConfig(dbConn *pgxpool.Pool, port, apiBase string) *AppConfig {
	dbQueries := data.New(dbConn)
	tmpl := template.Must(template.ParseGlob("./internal/web/templates/*"))

	return &AppConfig{DB: dbQueries, Templates: tmpl, Port: port, APIBase: apiBase}
}
