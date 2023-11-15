-- name: CreateAdmin :exec
INSERT INTO admin (email, hashed_password)
VALUES ($1, $2);

-- name: IsAdminByID :one
SELECT EXISTS(SELECT 1 FROM admin WHERE id = $1) AS "exists";

-- name: IsAdminByEmail :one
SELECT EXISTS(SELECT 1 FROM admin WHERE email = $1) AS "exists";

-- name: GetAdminByEmail :one
SELECT *
FROM admin
WHERE email = $1;

-- name: GetAdminEmailByID :one
SELECT email
FROM admin
WHERE id = $1;
