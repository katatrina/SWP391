-- name: IsUserExist :one
SELECT EXISTS(SELECT true FROM users WHERE id = $1);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: IsProviderRole :one
SELECT EXISTS (SELECT 1
               FROM users u
               WHERE u.id = $1
                 AND u.role_id = (SELECT id FROM roles WHERE name = 'provider')) AS is_provider;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: CreateCustomer :exec
INSERT INTO users (full_name, email, phone, address, role_id, hashed_password)
VALUES ($1, $2, $3, $4, 1, $5);

-- name: CreateProvider :one
INSERT INTO users (full_name, email, phone, address, role_id, hashed_password)
VALUES ($1, $2, $3, $4, 2, $5) RETURNING id;

-- name: CreateProviderDetails :exec
INSERT INTO providerDetails (user_id, company_name, tax_code)
VALUES ($1, $2, $3);