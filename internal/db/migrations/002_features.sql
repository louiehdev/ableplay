-- +goose Up
CREATE TABLE features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL CHECK (category IN ('Visual', 'Audio', 'Motor', 'Cognitive')),
    slug TEXT UNIQUE
);

-- +goose Down
DROP TABLE features;