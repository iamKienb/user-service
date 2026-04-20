-- name: InsertCredential :exec
INSERT INTO credentials (
    user_id,
    password_hash,
    password_version,
    password_updated_at
)
VALUES ($1, $2, $3, $4);

-- name: GetCredentialByID :one
SELECT *
FROM credentials
WHERE user_id = $1;