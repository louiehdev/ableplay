-- +goose Up
ALTER TABLE features
ADD COLUMN slug TEXT UNIQUE;

-- +goose Down
ALTER TABLE features
DROP COLUMN slug;