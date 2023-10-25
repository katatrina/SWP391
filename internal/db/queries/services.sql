-- name: CreateService :exec
INSERT INTO services (title, description, price, image_path, category_id, owned_by_provider_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: ListServiceByProvider :many
SELECT *
FROM services
WHERE owned_by_provider_id = $1;

-- name: GetServicesByCategorySlug :many
SELECT *
FROM services
WHERE category_id = (SELECT id FROM categories WHERE slug = $1);

-- name: GetServiceByID :one
SELECT *
FROM services
WHERE id = $1;

-- name: GetCompanyNameByServiceID :one
SELECT company_name
FROM provider_details
WHERE provider_id = (SELECT owned_by_provider_id FROM services WHERE services.id = $1);

-- name: GetServiceByCartItemID :one
SELECT *
FROM services
WHERE id = (SELECT service_id FROM cart_items WHERE cart_items.uuid = $1);