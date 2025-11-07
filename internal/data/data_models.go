package data

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type GamePublic struct {
	ID           string            `json:"id"`
	Title        string            `json:"title"`
	Developer    string            `json:"developer"`
	Publisher    string            `json:"publisher"`
	ReleaseYear  string            `json:"release_year"`
	Platforms    []string          `json:"platforms"`
	Description  string            `json:"description"`
	GameFeatures []GameFeatureData `json:"game_features"`
}

type GameData struct {
	ID          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Developer   pgtype.Text `json:"developer"`
	Publisher   pgtype.Text `json:"publisher"`
	ReleaseYear pgtype.Int4 `json:"release_year"`
	Platforms   []string    `json:"platforms"`
	Description pgtype.Text `json:"description"`
}

type FeaturePublic struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Category    string `json:"category"`
}

type FeatureData struct {
	ID          uuid.UUID   `json:"id"`
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
	Category    pgtype.Text `json:"category"`
}

type GameFeaturePublic struct {
	FeatureID string `json:"feature_id"`
	GameID    string `json:"game_id"`
	Name      string `json:"name"`
	Notes     string `json:"notes"`
	Verified  string `json:"verified"`
}

type GameFeatureData struct {
	FeatureID string `json:"id"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Notes     string `json:"notes"`
	Verified  bool   `json:"verified"`
}

type GameByFeaturePublic struct {
	FeatureID   string   `json:"feature_id"`
	GameID      string   `json:"game_id"`
	Notes       string   `json:"notes"`
	Verified    bool     `json:"verified"`
	Title       string   `json:"title"`
	Developer   string   `json:"developer"`
	Publisher   string   `json:"publisher"`
	ReleaseYear string   `json:"release_year"`
	Platforms   []string `json:"platforms"`
	Description string   `json:"description"`
}
