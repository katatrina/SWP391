// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: services.sql

package sqlc

import (
	"context"
)

const createService = `-- name: CreateService :exec
INSERT INTO services (title,
                      description,
                      price,
                      genre,
                      thumbnail_url,
                      category_id,
                      owned_by_user_id)
VALUES ($1, $2, $3, $4, $5, 1, $6)
`

type CreateServiceParams struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Price         int32  `json:"price"`
	Genre         string `json:"genre"`
	ThumbnailUrl  string `json:"thumbnail_url"`
	OwnedByUserID int32  `json:"owned_by_user_id"`
}

func (q *Queries) CreateService(ctx context.Context, arg CreateServiceParams) error {
	_, err := q.db.ExecContext(ctx, createService,
		arg.Title,
		arg.Description,
		arg.Price,
		arg.Genre,
		arg.ThumbnailUrl,
		arg.OwnedByUserID,
	)
	return err
}
