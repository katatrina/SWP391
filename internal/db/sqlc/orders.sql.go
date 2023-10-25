// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: orders.sql

package sqlc

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (uuid, buyer_id, seller_id, payment_method)
VALUES ($1, $2, $3, $4)
RETURNING uuid
`

type CreateOrderParams struct {
	UUID          string `json:"uuid"`
	BuyerID       int32  `json:"buyer_id"`
	SellerID      int32  `json:"seller_id"`
	PaymentMethod string `json:"payment_method"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.UUID,
		arg.BuyerID,
		arg.SellerID,
		arg.PaymentMethod,
	)
	var uuid string
	err := row.Scan(&uuid)
	return uuid, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (uuid, order_id, service_id, quantity, sub_total)
VALUES ($1, $2, $3, $4, $5)
RETURNING uuid
`

type CreateOrderItemParams struct {
	UUID      string `json:"uuid"`
	OrderID   string `json:"order_id"`
	ServiceID int32  `json:"service_id"`
	Quantity  int32  `json:"quantity"`
	SubTotal  int32  `json:"sub_total"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (string, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.UUID,
		arg.OrderID,
		arg.ServiceID,
		arg.Quantity,
		arg.SubTotal,
	)
	var uuid string
	err := row.Scan(&uuid)
	return uuid, err
}

const createOrderItemDetails = `-- name: CreateOrderItemDetails :exec
INSERT INTO order_item_details (order_item_id, title, price, image_path)
VALUES ($1, $2, $3, $4)
`

type CreateOrderItemDetailsParams struct {
	OrderItemID string `json:"order_item_id"`
	Title       string `json:"title"`
	Price       int32  `json:"price"`
	ImagePath   string `json:"image_path"`
}

func (q *Queries) CreateOrderItemDetails(ctx context.Context, arg CreateOrderItemDetailsParams) error {
	_, err := q.db.ExecContext(ctx, createOrderItemDetails,
		arg.OrderItemID,
		arg.Title,
		arg.Price,
		arg.ImagePath,
	)
	return err
}