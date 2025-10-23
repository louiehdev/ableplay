-- +goose Up
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    developer TEXT,
    publisher TEXT,
    release_year INT,
    platforms TEXT[],
    description TEXT
);

-- +goose Down
DROP TABLE games;