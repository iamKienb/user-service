-- name: SaveShopAddress :exec
INSERT INTO shop_addresses (
    id,
    shop_id, 

    country_id,
    city_id,
    district_id, 
    ward_id,

    address_line,
    contact_name,
    phone_number,
    type,

    created_at,
    updated_at,

    created_by,
    updated_by
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);

-- name: CheckRequiredAddresses :one
SELECT 
    EXISTS (
        SELECT 1 
        FROM shop_addresses sa_pickup 
        WHERE sa_pickup.shop_id = $1 AND sa_pickup.type = 'PICKUP'
    ) AS has_pickup,
    EXISTS (
        SELECT 1 
        FROM shop_addresses sa_return 
        WHERE sa_return.shop_id = $1 AND sa_return.type = 'RETURN'
    ) AS has_return;

