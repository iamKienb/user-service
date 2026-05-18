-- name: SaveShopProfile :exec
INSERT INTO shop_profiles (
    shop_id,
    description,
    logo_url,
    banner_url,
    created_by,
    updated_by,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetShopProfileByID :one
SELECT *
FROM shop_profiles
WHERE shop_id = $1;