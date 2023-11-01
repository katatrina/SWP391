// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: users.sql

package sqlc

import (
	"context"
	"time"
)

const createCustomer = `-- name: CreateCustomer :one
INSERT INTO users (full_name, email, phone, address, role_id, hashed_password)
VALUES ($1, $2, $3, $4, 1, $5) RETURNING id
`

type CreateCustomerParams struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"hashed_password"`
}

func (q *Queries) CreateCustomer(ctx context.Context, arg CreateCustomerParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createCustomer,
		arg.FullName,
		arg.Email,
		arg.Phone,
		arg.Address,
		arg.Password,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createProvider = `-- name: CreateProvider :one
INSERT INTO users (full_name, email, phone, address, role_id, hashed_password)
VALUES ($1, $2, $3, $4, 2, $5) RETURNING id
`

type CreateProviderParams struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Password string `json:"hashed_password"`
}

func (q *Queries) CreateProvider(ctx context.Context, arg CreateProviderParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createProvider,
		arg.FullName,
		arg.Email,
		arg.Phone,
		arg.Address,
		arg.Password,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createProviderDetails = `-- name: CreateProviderDetails :exec
INSERT INTO provider_details (provider_id, company_name, tax_code)
VALUES ($1, $2, $3)
`

type CreateProviderDetailsParams struct {
	ProviderID  int32  `json:"provider_id"`
	CompanyName string `json:"company_name"`
	TaxCode     string `json:"tax_code"`
}

func (q *Queries) CreateProviderDetails(ctx context.Context, arg CreateProviderDetailsParams) error {
	_, err := q.db.ExecContext(ctx, createProviderDetails, arg.ProviderID, arg.CompanyName, arg.TaxCode)
	return err
}

const getFullProviderInfo = `-- name: GetFullProviderInfo :one
SELECT u.id,
       u.full_name,
       u.email,
       u.phone,
       u.address,
       u.created_at,
       pd.company_name,
       pd.tax_code
FROM users u
         JOIN provider_details pd ON u.id = pd.provider_id
WHERE u.id = $1
`

type GetFullProviderInfoRow struct {
	ID          int32     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	CompanyName string    `json:"company_name"`
	TaxCode     string    `json:"tax_code"`
}

func (q *Queries) GetFullProviderInfo(ctx context.Context, id int32) (GetFullProviderInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getFullProviderInfo, id)
	var i GetFullProviderInfoRow
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.CreatedAt,
		&i.CompanyName,
		&i.TaxCode,
	)
	return i, err
}

const getProviderDetailsByID = `-- name: GetProviderDetailsByID :one
SELECT id, provider_id, company_name, tax_code, created_at
FROM provider_details
WHERE provider_id = $1
`

func (q *Queries) GetProviderDetailsByID(ctx context.Context, providerID int32) (ProviderDetail, error) {
	row := q.db.QueryRowContext(ctx, getProviderDetailsByID, providerID)
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, full_name, email, phone, address, role_id, hashed_password, created_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.RoleID,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, full_name, email, phone, address, role_id, hashed_password, created_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FullName,
		&i.Email,
		&i.Phone,
		&i.Address,
		&i.RoleID,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getUserRoleByID = `-- name: GetUserRoleByID :one
SELECT r.name
FROM users u
         JOIN roles r ON u.role_id = r.id
WHERE u.id = $1
`

func (q *Queries) GetUserRoleByID(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserRoleByID, id)
	var name string
	err := row.Scan(&name)
	return name, err
}

const isProvider = `-- name: IsProvider :one
SELECT EXISTS (SELECT 1
               FROM users u
               WHERE u.id = $1
                 AND u.role_id = (SELECT id FROM roles WHERE name = 'provider')) AS is_provider
`

func (q *Queries) IsProvider(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, isProvider, id)
	var is_provider bool
	err := row.Scan(&is_provider)
	return is_provider, err
}

const isUserExist = `-- name: IsUserExist :one
SELECT EXISTS(SELECT true FROM users WHERE id = $1)
`

func (q *Queries) IsUserExist(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, isUserExist, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const listProviders = `-- name: ListProviders :many
SELECT u.id,
       u.full_name,
       u.email,
       u.phone,
       u.address,
       u.created_at,
       pd.company_name,
       pd.tax_code
FROM users u
         JOIN provider_details pd ON u.id = pd.provider_id
WHERE u.role_id = (SELECT id FROM roles WHERE name = 'provider')
`

type ListProvidersRow struct {
	ID          int32     `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	CreatedAt   time.Time `json:"created_at"`
	CompanyName string    `json:"company_name"`
	TaxCode     string    `json:"tax_code"`
}

func (q *Queries) ListProviders(ctx context.Context) ([]ListProvidersRow, error) {
	rows, err := q.db.QueryContext(ctx, listProviders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListProvidersRow{}
	for rows.Next() {
		var i ListProvidersRow
		if err := rows.Scan(
			&i.ID,
			&i.FullName,
			&i.Email,
			&i.Phone,
			&i.Address,
			&i.CreatedAt,
			&i.CompanyName,
			&i.TaxCode,
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

const updateCustomerInfo = `-- name: UpdateCustomerInfo :exec
UPDATE users
SET full_name = $2,
    email     = $3,
    phone     = $4,
    address   = $5
WHERE id = $1
`

type UpdateCustomerInfoParams struct {
	ID       int32  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

func (q *Queries) UpdateCustomerInfo(ctx context.Context, arg UpdateCustomerInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateCustomerInfo,
		arg.ID,
		arg.FullName,
		arg.Email,
		arg.Phone,
		arg.Address,
	)
	return err
}

const updateProviderInfo = `-- name: UpdateProviderInfo :exec
UPDATE provider_details
SET company_name = $2,
    tax_code     = $3
WHERE provider_id = $1
`

type UpdateProviderInfoParams struct {
	ProviderID  int32  `json:"provider_id"`
	CompanyName string `json:"company_name"`
	TaxCode     string `json:"tax_code"`
}

func (q *Queries) UpdateProviderInfo(ctx context.Context, arg UpdateProviderInfoParams) error {
	_, err := q.db.ExecContext(ctx, updateProviderInfo, arg.ProviderID, arg.CompanyName, arg.TaxCode)
	return err
}
