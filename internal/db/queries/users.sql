-- name: IsUserExist :one
SELECT EXISTS(SELECT true FROM users WHERE id = $1);

-- name: CreateUser :exec
INSERT INTO users (name, email, hashed_password, created_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP);

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: UpdateUserPassword :exec
UPDATE users
SET hashed_password = $1
WHERE id = $2;
