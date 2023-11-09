-- name: CreateAdmin :exec
INSERT INTO admin (email, hashed_password)
VALUES ($1, $2);

-- name: IsAdmin :one
SELECT EXISTS(SELECT 1 FROM admin WHERE id = $1) AS "exists";