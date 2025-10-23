-- +goose Up
CREATE TABLE features (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    category TEXT CHECK (category IN ('Visual', 'Audio', 'Motor', 'Cognitive'))
);

-- +goose Down
DROP TABLE features;