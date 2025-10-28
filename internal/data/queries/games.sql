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
    title = COALESCE($2, title),
    developer = COALESCE($3, developer),
    publisher = COALESCE($4, publisher),
    release_year = COALESCE($5, release_year),
    platforms = COALESCE($6, platforms),
    description = COALESCE($7, description)
WHERE id = $1;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;