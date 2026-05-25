-- name: CreateUserCredential :exec
INSERT INTO user_credentials (
    user_id,
    password_hash,
    password_version,
    password_updated_at
)
VALUES ($1, $2, $3, $4);

-- name: FindUserCredentialByID :one
SELECT *
FROM user_credentials
WHERE user_id = $1;