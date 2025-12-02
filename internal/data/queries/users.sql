-- name: AddUser :exec
INSERT INTO users (created_at, updated_at, first_name, last_name, role, email, password)
VALUES (NOW(), NOW(), $1, $2, $3, $4, $5);

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

-- name: UpdateUser :exec
UPDATE users SET 
    updated_at = NOW(),
    first_name = $2,
    last_name = $3,
    role = $4,
    email = $5,
    password = $6
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;