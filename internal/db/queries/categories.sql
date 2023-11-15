-- name: ListCategories :many
SELECT *
FROM categories
ORDER BY id ASC;

-- name: ListCategoryIDs :many
SELECT id
FROM categories
ORDER BY id ASC;

-- name: IsCategoryExists :one
SELECT EXISTS(SELECT 1 FROM categories WHERE slug = $1);

-- name: GetServiceNumberByCategoryID :one
SELECT COUNT(*)
FROM services
WHERE category_id = $1 AND status = 'active';