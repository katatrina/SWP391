-- name: CreateService :exec
INSERT INTO services (title,
                      description,
                      price,
                      genre,
                      thumbnail_url,
                      category_id,
                      owned_by_user_id)
VALUES ($1, $2, $3, $4, $5, 1, $6);

-- name: ListServiceByProvider :many
SELECT *
FROM services
WHERE owned_by_user_id = $1;