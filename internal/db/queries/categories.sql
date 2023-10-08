-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY id ASC;

-- name: GetCategoryIDBySlug :one
SELECT id
FROM categories
WHERE slug = $1;