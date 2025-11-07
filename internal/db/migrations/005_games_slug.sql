-- +goose Up
ALTER TABLE games
ADD COLUMN slug TEXT UNIQUE;

-- +goose Down
ALTER TABLE games
DROP COLUMN slug;