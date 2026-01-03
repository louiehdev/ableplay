-- name: AddUser :one
INSERT INTO users (created_at, updated_at, first_name, last_name, email, password)
VALUES (NOW(), NOW(), $1, $2, $3, $4)
RETURNING *;

-- name: GetUsers :many
SELECT * FROM users
ORDER BY role DESC
LIMIT $1;

-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByAPIKey :one
SELECT users.id, users.created_at, users.updated_at, users.first_name, users.last_name, users.role, users.email, users.password FROM users
LEFT JOIN api_keys ON users.id = api_keys.user_id
WHERE api_keys.api_key = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users SET 
    updated_at = NOW(),
    first_name = $2,
    last_name = $3,
    email = $4,
    password = $5
WHERE id = $1
RETURNING *;

-- name: UpgradeUser :exec
UPDATE users SET
    updated_at = NOW(),
    role = $2
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;