-- name: AddGame :one
INSERT INTO games (created_at, updated_at, title, developer, publisher, release_year, platforms, description, slug)
VALUES (NOW(), NOW(), $1, $2, $3, $4, $5, $6, $7)
ON CONFLICT (slug) DO NOTHING
RETURNING *;

-- name: GetGamesWithFeatures :many
SELECT 
    games.id,
    games.title,
    games.developer,
    games.publisher,
    games.release_year,
    games.platforms,
    games.description,
    (
        SELECT json_agg(json_build_object('id', features.id, 'name', features.name, 'title', games.title, 'notes', games_features.notes, 'verified', games_features.verified))
        FROM games_features
        JOIN features ON features.id = games_features.feature_id
        WHERE games_features.game_id = games.id
    ) AS game_features
FROM games
ORDER BY games.title;


-- name: GetGames :many
SELECT id, title, developer, publisher, release_year, platforms, description
FROM games
ORDER BY title;

-- name: GetGame :one
SELECT id, title, developer, publisher, release_year, platforms, description FROM games
WHERE id = $1;

-- name: UpdateGame :exec
UPDATE games SET 
    updated_at = NOW(),
    title = $2,
    developer = $3,
    publisher = $4,
    release_year = $5,
    platforms = $6,
    description = $7
WHERE id = $1;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;