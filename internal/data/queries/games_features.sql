-- name: CreateGameFeature :one
INSERT INTO games_features (created_at, updated_at, game_id, feature_id, notes, verified)
VALUES (NOW(), NOW(), $1, $2, $3, $4)
RETURNING *;

-- name: GetGameFeature :one
SELECT
    games_features.notes, games_features.verified,
    games.title,
    features.name
FROM games_features
JOIN games ON games_features.game_id = games.id
JOIN features ON games_features.feature_id = features.id
WHERE games_features.game_id = $1 AND games_features.feature_id = $2;

-- name: GetFeaturesByGame :many
SELECT 
    games_features.notes, games_features.verified,
    features.id AS feature_id, features.name, features.description, features.category
FROM games_features
JOIN features ON games_features.feature_id = features.id
WHERE games_features.game_id = $1;

-- name: GetGamesByFeature :many
SELECT 
    games_features.notes, games_features.verified, games_features.feature_id,
    games.id AS game_id, games.title, games.developer, games.publisher, games.release_year, games.platforms, games.description
FROM games_features
JOIN games ON games_features.game_id = games.id
WHERE games_features.feature_id = $1;

-- name: UpdateGameFeature :exec
UPDATE games_features SET
    updated_at = NOW(),
    game_id = $3,
    feature_id = $4,
    notes = $5,
    verified = $6
WHERE game_id = $1 AND feature_id = $2;

-- name: DeleteGameFeature :exec
DELETE FROM games_features
WHERE game_id = $1 AND feature_id = $2;