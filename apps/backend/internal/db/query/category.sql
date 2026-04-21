-- name: ListCategories :many
SELECT * FROM categories ORDER BY created_at DESC;

-- name: GetCategoryById :one
SELECT * FROM categories WHERE id = $1 LIMIT 1;

-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;

-- name: UpdateCategory :one
UPDATE categories SET name = $1 WHERE id = $2 RETURNING *;

-- name: DeleteCategory :one
DELETE FROM categories WHERE id = $1 RETURNING *;

-- name: GetCategoryByName :one
SELECT * FROM categories WHERE name = $1 LIMIT 1;