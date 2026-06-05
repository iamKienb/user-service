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
VALUES (
    @id::uuid,
    @email::text,
    @email_verified_at::timestamptz,
    @status::text,
    @roles::text[],
    @created_at::timestamptz,
    @updated_at::timestamptz,
    @deleted_at::timestamptz
);

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
