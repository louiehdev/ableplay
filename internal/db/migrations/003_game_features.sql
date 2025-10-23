-- +goose Up
CREATE TABLE game_features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    game_id UUID NOT NULL,
    feature_id UUID UNIQUE NOT NULL,
    notes TEXT,
    verified BOOLEAN DEFAULT false,
    FOREIGN KEY(game_id) REFERENCES games (id) ON DELETE CASCADE,
    FOREIGN KEY(feature_id) REFERENCES features (id)
);

-- +goose Down
DROP TABLE game_features;