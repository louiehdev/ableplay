-- +goose Up
CREATE TABLE games (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    title TEXT NOT NULL,
    developer TEXT,
    publisher TEXT,
    release_year INT,
    platforms TEXT[],
    description TEXT,
    slug TEXT UNIQUE
);

-- +goose Down
DROP TABLE games;