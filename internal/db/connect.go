package db

import (
	"context"
	"embed"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	goose "github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrationsEmbed embed.FS

func Connect(ctx context.Context, dbURL string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dbURL)
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
