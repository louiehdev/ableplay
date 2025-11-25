-- +goose Up
CREATE TABLE games_changes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'denied')),
    user_id UUID NOT NULL REFERENCES users (id),
    moderator_id UUID REFERENCES users (id),
    title TEXT NOT NULL,
    developer TEXT,
    publisher TEXT,
    release_year INT,
    platforms TEXT[],
    description TEXT
);

CREATE TABLE features_changes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'denied')),
    user_id UUID NOT NULL REFERENCES users (id),
    moderator_id UUID REFERENCES users (id),
    name TEXT NOT NULL,
    description TEXT,
    category TEXT NOT NULL CHECK (category IN ('Visual', 'Audio', 'Motor', 'Cognitive'))
);

CREATE TABLE games_features_changes (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    status TEXT NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'denied')),
    user_id UUID NOT NULL REFERENCES users (id),
    moderator_id UUID REFERENCES users (id),
    game_id UUID NOT NULL REFERENCES games (id),
    feature_id UUID NOT NULL REFERENCES features (id),
    notes TEXT,
    verified BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE games_changes;
DROP TABLE features_changes;
DROP TABLE games_features_changes;