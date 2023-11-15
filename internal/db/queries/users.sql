-- name: IsUserExist :one
SELECT EXISTS(SELECT true FROM users WHERE id = $1);

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: IsProvider :one
SELECT EXISTS (SELECT 1
               FROM users u
               WHERE u.id = $1
                 AND u.role_id = (SELECT id FROM roles WHERE name = 'provider')) AS is_provider;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: CreateCustomer :one
INSERT INTO users (full_name, email, phone, address, role_id, hashed_password)
VALUES ($1, $2, $3, $4, 1, $5) RETURNING *;

-- name: CreateProvider :one
INSERT INTO users (full_name, email, phone, address, role_id, hashed_password)
VALUES ($1, $2, $3, $4, 2, $5) RETURNING id;

-- name: CreateProviderDetails :exec
INSERT INTO provider_details (provider_id, company_name, tax_code)
VALUES ($1, $2, $3);

-- name: ListProviders :many
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
WHERE u.role_id = (SELECT id FROM roles WHERE name = 'provider');

-- name: GetFullProviderInfo :one
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
WHERE u.id = $1;

-- name: GetUserRoleByID :one
SELECT r.name
FROM users u
         JOIN roles r ON u.role_id = r.id
WHERE u.id = $1;

-- name: GetProviderDetailsByID :one
SELECT *
FROM provider_details
WHERE provider_id = $1;

-- name: UpdateCustomerInfo :exec
UPDATE users
SET full_name = $2,
    email     = $3,
    phone     = $4,
    address   = $5
WHERE id = $1;

-- name: UpdateProviderInfo :exec
UPDATE provider_details
SET company_name = $2,
    tax_code     = $3
WHERE provider_id = $1;

-- name: GetCustomerNumber :one
SELECT COUNT(*)
FROM users
WHERE role_id = (SELECT id FROM roles WHERE name = 'customer');

-- name: GetProviderNumber :one
SELECT COUNT(*)
FROM users
WHERE role_id = (SELECT id FROM roles WHERE name = 'provider');

-- name: ListCustomers :many
SELECT *
FROM users
WHERE role_id = (SELECT id FROM roles WHERE name = 'customer')
ORDER BY created_at DESC;

-- name: GetProviders :many
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
ORDER BY u.created_at DESC;

-- name: DeleteAccount :one
DELETE
FROM users
WHERE id = $1
RETURNING full_name;