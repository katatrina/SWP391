-- name: CreateCart :exec
INSERT INTO carts (user_id)
VALUES ($1);

-- name: AddServiceToCart :exec
INSERT INTO cart_items (cart_id, service_id, quantity, price)
VALUES ($1, $2, $3, $4);

-- name: IsServiceExists :one
SELECT EXISTS(SELECT 1 FROM cart_items WHERE cart_id = $1 AND service_id = $2);

-- name: GetCartIDByUserId :one
SELECT id FROM carts WHERE user_id = $1;