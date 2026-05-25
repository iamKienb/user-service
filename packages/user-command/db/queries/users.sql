-- name: CreateUser :exec
INSERT INTO users (
    id,
    email,
    email_verified_at,
    status,
    roles,
    created_at,
    updated_at,
    deleted_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: FindUserByID :one
SELECT * 
FROM users
WHERE id = $1;

-- name: FindUserByEmail :one
SELECT * 
FROM users
WHERE email = $1;

-- name: UpdateUser :exec
UPDATE users
SET
    email = $2,
    email_verified_at = $3,
    status = $4,
    roles = $5,
    updated_at = $6,
    deleted_at = $7
WHERE id = $1;
