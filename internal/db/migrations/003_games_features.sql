-- +goose Up
CREATE TABLE games_features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    game_id UUID NOT NULL,
    feature_id UUID NOT NULL,
    notes TEXT,
    verified BOOLEAN NOT NULL DEFAULT false,
    FOREIGN KEY(game_id) REFERENCES games (id) ON DELETE CASCADE,
    FOREIGN KEY(feature_id) REFERENCES features (id)
);

-- +goose Down
DROP TABLE games_features;