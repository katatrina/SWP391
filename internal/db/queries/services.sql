-- name: CreateService :exec
INSERT INTO services (title,
                      description,
                      price,
                      category_id,
                      thumbnail_url,
                      owned_by_user_id)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: ListServiceByProvider :many
SELECT *
FROM services
WHERE owned_by_user_id = $1;