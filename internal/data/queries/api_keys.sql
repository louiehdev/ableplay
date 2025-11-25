-- name: CreateKey :one
INSERT INTO api_keys (id, api_key, created_at, user_id)
VALUES ($1, $2, NOW(), $3)
RETURNING *;

-- name: GetKeyByUser :one
SELECT api_key, created_at
FROM api_keys
WHERE user_id = $1;

-- name: GetKeyByID :one
SELECT api_key
FROM api_keys
WHERE id = $1;

-- name: UpdateKey :exec
UPDATE api_keys SET
    id = $2,
    api_key = $3,
    created_at = NOW()
WHERE user_id = $1;

-- name: DeleteKey :exec
DELETE FROM api_keys
WHERE user_id = $1;