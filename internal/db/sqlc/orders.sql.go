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
VALUES ($1, $2, $3, $4) RETURNING uuid, buyer_id, seller_id, status, payment_method, grand_total, created_at
`

type CreateOrderParams struct {
	UUID          string `json:"uuid"`
	BuyerID       int32  `json:"buyer_id"`
	SellerID      int32  `json:"seller_id"`
	PaymentMethod string `json:"payment_method"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRowContext(ctx, createOrder,
		arg.UUID,
		arg.BuyerID,
		arg.SellerID,
		arg.PaymentMethod,
	)
	var i Order
	err := row.Scan(
		&i.UUID,
		&i.BuyerID,
		&i.SellerID,
		&i.Status,
		&i.PaymentMethod,
		&i.GrandTotal,
		&i.CreatedAt,
	)
	return i, err
}

const createOrderItem = `-- name: CreateOrderItem :one
INSERT INTO order_items (uuid, order_id, service_id, quantity, sub_total)
VALUES ($1, $2, $3, $4, $5) RETURNING uuid, order_id, service_id, quantity, sub_total, created_at
`

type CreateOrderItemParams struct {
	UUID      string `json:"uuid"`
	OrderID   string `json:"order_id"`
	ServiceID int32  `json:"service_id"`
	Quantity  int32  `json:"quantity"`
	SubTotal  int32  `json:"sub_total"`
}

func (q *Queries) CreateOrderItem(ctx context.Context, arg CreateOrderItemParams) (OrderItem, error) {
	row := q.db.QueryRowContext(ctx, createOrderItem,
		arg.UUID,
		arg.OrderID,
		arg.ServiceID,
		arg.Quantity,
		arg.SubTotal,
	)
	var i OrderItem
	err := row.Scan(
		&i.UUID,
		&i.OrderID,
		&i.ServiceID,
		&i.Quantity,
		&i.SubTotal,
		&i.CreatedAt,
	)
	return i, err
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

const getFullOrderItemsInformationByOrderId = `-- name: GetFullOrderItemsInformationByOrderId :many
SELECT oi.uuid,
       oi.order_id,
       oi.service_id,
       oi.quantity,
       oid.title,
       oid.image_path,
       oid.price
FROM order_items oi
         INNER JOIN order_item_details oid ON oid.order_item_id = oi.uuid
WHERE oi.order_id = $1
`

type GetFullOrderItemsInformationByOrderIdRow struct {
	UUID      string `json:"uuid"`
	OrderID   string `json:"order_id"`
	ServiceID int32  `json:"service_id"`
	Quantity  int32  `json:"quantity"`
	Title     string `json:"title"`
	ImagePath string `json:"image_path"`
	Price     int32  `json:"price"`
}

func (q *Queries) GetFullOrderItemsInformationByOrderId(ctx context.Context, orderID string) ([]GetFullOrderItemsInformationByOrderIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getFullOrderItemsInformationByOrderId, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFullOrderItemsInformationByOrderIdRow{}
	for rows.Next() {
		var i GetFullOrderItemsInformationByOrderIdRow
		if err := rows.Scan(
			&i.UUID,
			&i.OrderID,
			&i.ServiceID,
			&i.Quantity,
			&i.Title,
			&i.ImagePath,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrderByOrderItemID = `-- name: GetOrderByOrderItemID :one
SELECT uuid, buyer_id, seller_id, status, payment_method, grand_total, created_at
FROM orders
WHERE uuid = (SELECT order_id
              FROM order_items
              WHERE order_items.uuid = $1)
`

func (q *Queries) GetOrderByOrderItemID(ctx context.Context, uuid string) (Order, error) {
	row := q.db.QueryRowContext(ctx, getOrderByOrderItemID, uuid)
	var i Order
	err := row.Scan(
		&i.UUID,
		&i.BuyerID,
		&i.SellerID,
		&i.Status,
		&i.PaymentMethod,
		&i.GrandTotal,
		&i.CreatedAt,
	)
	return i, err
}

const getPurchaseOrders = `-- name: GetPurchaseOrders :many
SELECT uuid, buyer_id, seller_id, status, payment_method, grand_total, created_at
FROM orders
WHERE buyer_id = $1
ORDER BY created_at DESC
`

func (q *Queries) GetPurchaseOrders(ctx context.Context, buyerID int32) ([]Order, error) {
	rows, err := q.db.QueryContext(ctx, getPurchaseOrders, buyerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.UUID,
			&i.BuyerID,
			&i.SellerID,
			&i.Status,
			&i.PaymentMethod,
			&i.GrandTotal,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrderTotal = `-- name: UpdateOrderTotal :exec
UPDATE orders
SET grand_total = $1
WHERE uuid = $2
`

type UpdateOrderTotalParams struct {
	GrandTotal int32  `json:"grand_total"`
	UUID       string `json:"uuid"`
}

func (q *Queries) UpdateOrderTotal(ctx context.Context, arg UpdateOrderTotalParams) error {
	_, err := q.db.ExecContext(ctx, updateOrderTotal, arg.GrandTotal, arg.UUID)
	return err
}
