package db

import (
	"context"
	"embed"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	goose "github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsEmbed embed.FS

func Connect(ctx context.Context) (*pgxpool.Pool, error) {
	connStr := os.Getenv("DB_URL")

	config, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(pool *pgxpool.Pool) error {
	goose.SetBaseFS(migrationsEmbed)
	db := stdlib.OpenDBFromPool(pool)
	defer db.Close()
	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}
	return nil
}
