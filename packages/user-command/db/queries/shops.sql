-- name: SaveShop :exec
INSERT INTO shops (
    id,
    owner_id,
    name,
    slug,
    status,
    created_by,
    updated_by,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: UpdateShopStatus :exec
UPDATE shops
SET
    status = $2,
    updated_by = $3,
    updated_at = $4
WHERE id = $1;


-- name: GetShopByID :one
SELECT *
FROM shops
WHERE id = $1;

-- name: CountBySlug :one
SELECT 
    1
FROM shops 
WHERE slug = $1;