-- name: CreateCart :exec
INSERT INTO carts (user_id)
VALUES ($1);

-- name: AddServiceToCart :exec
INSERT INTO cart_items (cart_id, service_id, quantity, sub_total)
VALUES ($1, $2, $3, $4);

-- name: IsServiceAlreadyInCart :one
SELECT EXISTS(SELECT 1
              FROM cart_items
              WHERE cart_id = $1
                AND service_id = $2);


-- name: GetCartIDByUserId :one
SELECT id
FROM carts
WHERE user_id = $1;

-- name: GetCartItemsByCartID :many
SELECT cart_items.id,
       cart_items.cart_id,
       cart_items.service_id,
       cart_items.quantity,
       cart_items.sub_total,
       services.title,
       services.price,
       services.image_path,
       services.owned_by_provider_id
FROM cart_items
         INNER JOIN services ON cart_items.service_id = services.id
WHERE cart_items.cart_id = $1;

-- name: GetCartItemByCartIDAndServiceID :one
SELECT *
FROM cart_items
WHERE cart_id = $1
  AND service_id = $2;

-- name: UpdateCartItemQuantityAndSubTotal :exec
UPDATE cart_items
SET quantity  = $1,
    sub_total = $2
WHERE id = $3;

-- name: CreateCartItem :exec
INSERT INTO cart_items (cart_id, service_id, quantity, sub_total)
VALUES ($1, $2, $3, $4);
