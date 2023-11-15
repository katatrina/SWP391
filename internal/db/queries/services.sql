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
WHERE category_id = (SELECT id FROM categories WHERE slug = $1)
  AND owned_by_provider_id != $2;

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

-- name: GetProviderDetailsByServiceID :one
SELECT provider_details.company_name, u.address, u.phone, u.email
FROM provider_details
         INNER JOIN users u on provider_details.provider_id = u.id
WHERE provider_details.provider_id = (SELECT owned_by_provider_id FROM services WHERE services.id = $1);

-- name: IsUserUsedService :one
SELECT EXISTS(SELECT 1
              FROM orders
                       INNER JOIN order_items o_i on orders.uuid = o_i.order_id
              WHERE orders.buyer_id = $1
                AND o_i.service_id = $2
                AND orders.status_id = (SELECT id FROM order_status WHERE code = 'completed'));

-- name: GetServiceNumber :one
SELECT COUNT(*)
FROM services
WHERE status = 'active';

-- name: ListServices :many
SELECT *
FROM services
WHERE status = 'inactive'
  and owned_by_provider_id != $1
ORDER BY created_at DESC;

-- name: ListInactiveServices :many
SELECT *
FROM services
WHERE status = 'inactive'
ORDER BY created_at DESC;

-- name: GetCategoryNameByServiceID :one
SELECT name
FROM categories
WHERE id = (SELECT category_id FROM services WHERE services.id = $1);

-- name: UpdateServiceStatus :exec
UPDATE services
SET status = $1,
    reject_reason = $2
WHERE id = $3;