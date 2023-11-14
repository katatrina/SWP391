-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY id ASC;

-- name: ListCategoryIDs :many
SELECT id
FROM categories
ORDER BY id ASC;

-- name: GetCategoryBySlug :one
SELECT *
FROM categories
WHERE slug = $1 LIMIT 1;

-- name: GetServiceNumberByCategoryID :one
SELECT COUNT(*)
FROM services
WHERE category_id = $1;