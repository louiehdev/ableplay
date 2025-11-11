-- name: AddFeature :one
INSERT INTO features (created_at, updated_at, name, description, category, slug)
VALUES (NOW(), NOW(), $1, $2, $3, LOWER(REPLACE(CONCAT($1::text, '-', $3::text), ' ', '-')))
ON CONFLICT (slug) DO NOTHING
RETURNING *;

-- name: GetFeaturesSearch :many
SELECT id, name, description, category 
FROM features
WHERE 
    name ILIKE CONCAT('%', $1::text, '%')
    OR category ILIKE CONCAT('%', $1::text, '%')
ORDER BY category DESC;

-- name: GetFeatures :many
SELECT id, name, description, category 
FROM features
ORDER BY category DESC
LIMIT $1;

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