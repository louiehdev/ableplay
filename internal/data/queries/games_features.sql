-- name: CreateGameFeature :one
INSERT INTO games_features (created_at, updated_at, game_id, feature_id, notes, verified)
VALUES (NOW(), NOW(), $1, $2, $3, $4)
RETURNING *;

-- name: GetFeaturesByGame :many
SELECT 
    games_features.id, games_features.notes, games_features.verified,
    features.id AS feature_id, features.name, features.description, features.category
FROM games_features
INNER JOIN features ON games_features.feature_id = features.id
WHERE games_features.game_id = $1;

-- name: GetGamesByFeature :many
SELECT 
    games_features.id, games_features.notes, games_features.verified,
    games.id AS game_id, games.title, games.developer, games.publisher, games.release_year, games.platforms, games.description
FROM games_features
INNER JOIN games ON games_features.game_id = games.id
WHERE games_features.feature_id = $1;

-- name: DeleteGameFeature :exec
DELETE FROM games_features
WHERE id = $1;