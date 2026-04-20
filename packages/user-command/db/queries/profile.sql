-- name: InsertProfile :exec
INSERT INTO profiles (
    user_id,
    full_name,
    gender,
    phone_number,
    avatar_url,
    date_of_birth,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetProfileByID :one
SELECT *
FROM profiles
WHERE user_id = $1;