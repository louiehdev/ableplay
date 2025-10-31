-- name: AddGame :one
INSERT INTO games (created_at, updated_at, title, developer, publisher, release_year, platforms, description)
VALUES (NOW(), NOW(), $1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetGames :many
SELECT id, title, developer, publisher, release_year, platforms, description FROM games;

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