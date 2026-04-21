-- name: ListProducts :many
SELECT * FROM products ORDER BY created_at DESC;

-- name: GetProductById :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: CreateProduct :one
INSERT INTO
    products (
        product_image_file,
        name,
        price,
        stock,
        category_id
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING
    *;

-- name: UpdateProduct :one
UPDATE products
SET
    product_image_file = COALESCE($1, product_image_file),
    name = COALESCE($2, name),
    price = COALESCE($3, price),
    stock = COALESCE($4, stock),
    category_id = COALESCE($5, category_id)
WHERE
    id = $6
RETURNING
    *;

-- name: DeleteProduct :one
DELETE FROM products WHERE id = $1 RETURNING *;

-- name: GetProductByName :one
SELECT * FROM products WHERE name = $1 LIMIT 1;