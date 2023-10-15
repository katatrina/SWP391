-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY id ASC;

-- name: ListCategoryIDs :many
SELECT id
FROM categories
ORDER BY id ASC;