-- name: CreateOrder :one
INSERT INTO orders (uuid, buyer_id, seller_id, payment_method)
VALUES ($1, $2, $3, $4)
RETURNING uuid;

-- name: CreateOrderItem :one
INSERT INTO order_items (uuid, order_id, service_id, quantity, sub_total)
VALUES ($1, $2, $3, $4, $5)
RETURNING uuid;

-- name: CreateOrderItemDetails :exec
INSERT INTO order_item_details (order_item_id, title, price, image_path)
VALUES ($1, $2, $3, $4);