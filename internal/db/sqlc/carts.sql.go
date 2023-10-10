// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: carts.sql

package sqlc

import (
	"context"
)

const addServiceToCart = `-- name: AddServiceToCart :exec
INSERT INTO cart_items (cart_id, service_id, quantity, price)
VALUES ($1, $2, $3, $4)
`

type AddServiceToCartParams struct {
	CartID    int32 `json:"cart_id"`
	ServiceID int32 `json:"service_id"`
	Quantity  int32 `json:"quantity"`
	Price     int32 `json:"price"`
}

func (q *Queries) AddServiceToCart(ctx context.Context, arg AddServiceToCartParams) error {
	_, err := q.db.ExecContext(ctx, addServiceToCart,
		arg.CartID,
		arg.ServiceID,
		arg.Quantity,
		arg.Price,
	)
	return err
}

const createCart = `-- name: CreateCart :exec
INSERT INTO carts (user_id)
VALUES ($1)
`

func (q *Queries) CreateCart(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, createCart, userID)
	return err
}

const getCartIDByUserId = `-- name: GetCartIDByUserId :one
SELECT id FROM carts WHERE user_id = $1
`

func (q *Queries) GetCartIDByUserId(ctx context.Context, userID int32) (int32, error) {
	row := q.db.QueryRowContext(ctx, getCartIDByUserId, userID)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const isServiceExists = `-- name: IsServiceExists :one
SELECT EXISTS(SELECT 1 FROM cart_items WHERE cart_id = $1 AND service_id = $2)
`

type IsServiceExistsParams struct {
	CartID    int32 `json:"cart_id"`
	ServiceID int32 `json:"service_id"`
}

func (q *Queries) IsServiceExists(ctx context.Context, arg IsServiceExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, isServiceExists, arg.CartID, arg.ServiceID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
