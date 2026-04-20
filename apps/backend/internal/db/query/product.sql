-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: CreateProduct :one
INSERT INTO
    products (
        name,
        price,
        stock,
        category_id
    )
VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- name: UpdateProduct :one
UPDATE products
SET
    name = $1,
    price = $2,
    stock = $3,
    category_id = $4
WHERE
    id = $5
RETURNING
    *;

-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1 RETURNING *;