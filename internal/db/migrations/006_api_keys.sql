-- +goose Up
CREATE TABLE api_keys (
    id TEXT PRIMARY KEY,
    api_key TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    expires_at TIMESTAMP,
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE api_keys;