-- name: AddGame :one
INSERT INTO games (created_at, updated_at, title, developer, publisher, release_year, platforms, description)
VALUES (NOW(), NOW(), $1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetGames :many
SELECT title, developer, publisher, release_year, platforms, description FROM games;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;