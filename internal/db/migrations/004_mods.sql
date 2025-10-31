-- +goose Up
CREATE TABLE mods (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    game_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    source_url TEXT NOT NULL,
    verified BOOLEAN NOT NULL DEFAULT false,
    FOREIGN KEY(game_id) REFERENCES games (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE mods;