-- name: AddGamesChange :exec
INSERT INTO games_changes (id, user_id, change_type, title, developer, publisher, release_year, platforms, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: SubmitGamesChange :exec
INSERT INTO games (id, title, developer, publisher, release_year, platforms, description, slug)
SELECT 
    games_changes.id, 
    games_changes.title, 
    games_changes.developer, 
    games_changes.publisher, 
    games_changes.release_year, 
    games_changes.platforms, 
    games_changes.description, 
    LOWER(REPLACE(CONCAT(games_changes.title, '-', games_changes.developer), ' ', '-'))
FROM games_changes
WHERE games_changes.id = $1
ON CONFLICT (slug) DO NOTHING;

-- name: GetGamesChanges :many
SELECT * FROM games_changes
ORDER BY title
LIMIT $1;

-- name: GetGamesChangeByID :one
SELECT * FROM games_changes
WHERE id = $1;

-- name: UpdateGamesChange :exec
UPDATE games_changes SET 
    status = $2,
    moderator_id = $3
WHERE id = $1;

-- name: AddFeaturesChange :exec
INSERT INTO features_changes (id, user_id, change_type, name, description, category)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: SubmitFeaturesChange :exec
INSERT INTO features (id, name, description, category, slug)
SELECT 
    features_changes.id, 
    features_changes.name, 
    features_changes.description, 
    features_changes.category, 
    LOWER(REPLACE(CONCAT(features_changes.name, '-', features_changes.category), ' ', '-'))
FROM features_changes
WHERE features_changes.id = $1
ON CONFLICT (slug) DO NOTHING;

-- name: GetFeaturesChanges :many
SELECT * FROM features_changes
ORDER BY category DESC
LIMIT $1;

-- name: GetFeaturesChangeByID :one
SELECT * FROM features_changes
WHERE id = $1;

-- name: UpdateFeaturesChange :exec
UPDATE features_changes SET 
    status = $2,
    moderator_id = $3
WHERE id = $1;

-- name: AddGamesFeaturesChange :exec
INSERT INTO games_features_changes (id, user_id, change_type, game_id, feature_id, notes, verified)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: SubmitGamesFeaturesChange :exec
INSERT INTO games_features (id, updated_at, game_id, feature_id, notes, verified)
SELECT 
    games_features_changes.id,
    NOW(),
    games_features_changes.game_id,
    games_features_changes.feature_id, 
    games_features_changes.notes, 
    games_features_changes.verified
FROM games_features_changes
WHERE games_features_changes.id = $1;

-- name: GetGamesFeaturesChanges :many
SELECT * FROM games_features_changes
LIMIT $1;

-- name: GetGamesFeaturesChangeByID :one
SELECT * FROM games_features_changes
WHERE id = $1;

-- name: UpdateGamesFeaturesChange :exec
UPDATE games_features_changes SET 
    status = $2,
    moderator_id = $3
WHERE id = $1;