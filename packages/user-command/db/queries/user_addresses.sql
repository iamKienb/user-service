-- name: CreateUserAddress :exec
INSERT INTO user_addresses (
    id,
    user_id, 
    country_id, 
    city_id, 
    district_id, 
    ward_id, 
    address_line, 
    receiver_name, 
    phone_number, 
    label, 
    is_default,
    created_at, 
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13);