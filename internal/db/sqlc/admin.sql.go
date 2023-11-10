// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: admin.sql

package sqlc

import (
	"context"
)

const createAdmin = `-- name: CreateAdmin :exec
INSERT INTO admin (email, hashed_password)
VALUES ($1, $2)
`

type CreateAdminParams struct {
	Email    string `json:"email"`
	Password string `json:"hashed_password"`
}

func (q *Queries) CreateAdmin(ctx context.Context, arg CreateAdminParams) error {
	_, err := q.db.ExecContext(ctx, createAdmin, arg.Email, arg.Password)
	return err
}

const getAdminByEmail = `-- name: GetAdminByEmail :one
SELECT id, email, hashed_password
FROM admin
WHERE email = $1
`

func (q *Queries) GetAdminByEmail(ctx context.Context, email string) (Admin, error) {
	row := q.db.QueryRowContext(ctx, getAdminByEmail, email)
	var i Admin
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return i, err
}

const isAdminByEmail = `-- name: IsAdminByEmail :one
SELECT EXISTS(SELECT 1 FROM admin WHERE email = $1) AS "exists"
`

func (q *Queries) IsAdminByEmail(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRowContext(ctx, isAdminByEmail, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const isAdminByID = `-- name: IsAdminByID :one
SELECT EXISTS(SELECT 1 FROM admin WHERE id = $1) AS "exists"
`

func (q *Queries) IsAdminByID(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRowContext(ctx, isAdminByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
