-- name: AddFeature :one
INSERT INTO features (created_at, updated_at, name, description, category)
VALUES (NOW(), NOW(), $1, $2, $3)
RETURNING *;

-- name: GetFeatures :many
SELECT id, name, description, category FROM features;

-- name: GetFeature :one
SELECT id, name, description, category FROM features
WHERE id = $1;

-- name: UpdateFeature :exec
UPDATE features SET 
    updated_at = NOW(),
    name = $2,
    description = $3,
    category = $4
WHERE id = $1;

-- name: DeleteFeature :exec
DELETE FROM features
WHERE id = $1;