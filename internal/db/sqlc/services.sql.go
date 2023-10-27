// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: services.sql

package sqlc

import (
	"context"
)

const createService = `-- name: CreateService :exec
INSERT INTO services (title, description, price, image_path, category_id, owned_by_provider_id)
VALUES ($1, $2, $3, $4, $5, $6)
`

type CreateServiceParams struct {
	Title             string `json:"title"`
	Description       string `json:"description"`
	Price             int32  `json:"price"`
	ImagePath         string `json:"image_path"`
	CategoryID        int32  `json:"category_id"`
	OwnedByProviderID int32  `json:"owned_by_provider_id"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) error {
	_, err := q.db.ExecContext(ctx, createService,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.ImagePath,
		arg.CategoryID,
		arg.OwnedByProviderID,
	)
	return err
}

const getCompanyNameByServiceID = `-- name: GetCompanyNameByServiceID :one
SELECT company_name
FROM provider_details
WHERE provider_id = (SELECT owned_by_provider_id FROM services WHERE services.id = $1)
`

func (q *Queries) GetCompanyNameByServiceID(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, getCompanyNameByServiceID, id)
	var company_name string
	err := row.Scan(&company_name)
	return company_name, err
}

const getProviderDetailsByServiceID = `-- name: GetProviderDetailsByServiceID :one
SELECT id, provider_id, company_name, tax_code, created_at
FROM provider_details
WHERE provider_id = (SELECT owned_by_provider_id FROM services WHERE services.id = $1)
`

func (q *Queries) GetProviderDetailsByServiceID(ctx context.Context, id int32) (ProviderDetail, error) {
	row := q.db.QueryRowContext(ctx, getProviderDetailsByServiceID, id)
	var i ProviderDetail
	err := row.Scan(
		&i.ID,
		&i.ProviderID,
		&i.CompanyName,
		&i.TaxCode,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceByCartItemID = `-- name: GetServiceByCartItemID :one
SELECT id, title, description, price, image_path, category_id, owned_by_provider_id, status, created_at
FROM services
WHERE id = (SELECT service_id FROM cart_items WHERE cart_items.uuid = $1)
`

func (q *Queries) GetServiceByCartItemID(ctx context.Context, uuid string) (Service, error) {
	row := q.db.QueryRowContext(ctx, getServiceByCartItemID, uuid)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.ImagePath,
		&i.CategoryID,
		&i.OwnedByProviderID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getServiceByID = `-- name: GetServiceByID :one
SELECT id, title, description, price, image_path, category_id, owned_by_provider_id, status, created_at
FROM services
WHERE id = $1
`

func (q *Queries) GetServiceByID(ctx context.Context, id int32) (Service, error) {
	row := q.db.QueryRowContext(ctx, getServiceByID, id)
	var i Service
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Price,
		&i.ImagePath,
		&i.CategoryID,
		&i.OwnedByProviderID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getServicesByCategorySlug = `-- name: GetServicesByCategorySlug :many
SELECT id, title, description, price, image_path, category_id, owned_by_provider_id, status, created_at
FROM services
WHERE category_id = (SELECT id FROM categories WHERE slug = $1) AND owned_by_provider_id != $2
`

type GetServicesByCategorySlugParams struct {
	Slug              string `json:"slug"`
	OwnedByProviderID int32  `json:"owned_by_provider_id"`
}

func (q *Queries) GetServicesByCategorySlug(ctx context.Context, arg GetServicesByCategorySlugParams) ([]Service, error) {
	rows, err := q.db.QueryContext(ctx, getServicesByCategorySlug, arg.Slug, arg.OwnedByProviderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Price,
			&i.ImagePath,
			&i.CategoryID,
			&i.OwnedByProviderID,
			&i.Status,
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

const listServiceByProvider = `-- name: ListServiceByProvider :many
SELECT id, title, description, price, image_path, category_id, owned_by_provider_id, status, created_at
FROM services
WHERE owned_by_provider_id = $1
`

func (q *Queries) ListServiceByProvider(ctx context.Context, ownedByProviderID int32) ([]Service, error) {
	rows, err := q.db.QueryContext(ctx, listServiceByProvider, ownedByProviderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Service{}
	for rows.Next() {
		var i Service
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Price,
			&i.ImagePath,
			&i.CategoryID,
			&i.OwnedByProviderID,
			&i.Status,
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
