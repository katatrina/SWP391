-- name: CreateOrder :one
INSERT INTO orders (uuid, buyer_id, seller_id, status_id, payment_method)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: CreateOrderItem :one
INSERT INTO order_items (uuid, order_id, service_id, quantity, sub_total)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: CreateOrderItemDetails :exec
INSERT INTO order_item_details (order_item_id, title, price, image_path)
VALUES ($1, $2, $3, $4);

-- name: GetPurchaseOrders :many
SELECT o.uuid,
       o.buyer_id,
       o.seller_id,
       o.status_id,
       o.payment_method,
       o.grand_total,
       o.created_at,
       os.id,
       os.code,
       os.detail as status_detail
FROM orders AS o
         INNER JOIN order_status os ON os.id = o.status_id
WHERE buyer_id = $1
ORDER BY created_at;

-- name: GetPurchaseOrdersWithStatusCode :many
SELECT o.uuid,
       o.buyer_id,
       o.seller_id,
       o.status_id,
       o.payment_method,
       o.grand_total,
       o.created_at,
       os.id,
       os.code,
       os.detail as status_detail
FROM orders AS o
         INNER JOIN order_status os ON os.id = o.status_id
WHERE o.buyer_id = $1
  AND os.code = $2
ORDER BY created_at;

-- name: GetSellOrdersWithStatusCode :many
SELECT o.uuid,
       o.buyer_id,
       o.seller_id,
       o.status_id,
       o.payment_method,
       o.grand_total,
       o.created_at,
       os.id,
       os.code,
       os.detail as status_detail
FROM orders AS o
         INNER JOIN order_status os ON os.id = o.status_id
WHERE o.seller_id = $1
  AND os.code = $2
ORDER BY created_at;

-- name: UpdateOrderTotal :exec
UPDATE orders
SET grand_total = $1
WHERE uuid = $2;

-- name: GetFullOrderItemsInformationByOrderId :many
SELECT oi.uuid,
       oi.order_id,
       oi.service_id,
       oi.quantity,
       oid.title,
       oid.image_path,
       oid.price
FROM order_items oi
         INNER JOIN order_item_details oid ON oid.order_item_id = oi.uuid
WHERE oi.order_id = $1;

-- name: GetSellOrders :many
SELECT o.uuid,
       o.buyer_id,
       o.seller_id,
       o.status_id,
       o.payment_method,
       o.grand_total,
       o.created_at,
       os.id,
       os.code,
       os.detail as status_detail
FROM orders AS o
         INNER JOIN order_status os ON os.id = o.status_id
WHERE seller_id = $1
ORDER BY created_at;

-- name: GetOrderByOrderItemID :one
SELECT *
FROM orders
WHERE uuid = (SELECT order_id
              FROM order_items
              WHERE order_items.uuid = $1);


-- name: GetOrderStatuses :many
SELECT *
FROM order_status
ORDER BY id ASC;

-- name: UpdateOrderStatus :exec
UPDATE orders
SET status_id = (SELECT id
                 FROM order_status
                 WHERE code = $1)
WHERE uuid = $2;