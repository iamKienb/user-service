-- name: CreateUserAddress :exec
INSERT INTO user_addresses (
    id,
    user_id, 
    country_id,
    country_name,
    province_id,
    province_name,
    ward_id, 
    ward_name,
    address_line, 
    receiver_name, 
    phone_number, 
    label, 
    is_default,
    created_at, 
    updated_at
)
VALUES (
    @id::uuid,
    @user_id::uuid,
    @country_id::text,
    @country_name::text,
    @province_id::text,
    @province_name::text,
    @ward_id::text,
    @ward_name::text,
    @address_line::text,
    @receiver_name::text,
    @phone_number::text,
    @label::text,
    @is_default::boolean,
    @created_at::timestamptz,
    @updated_at::updated_at
);

-- name: FindUserAddressByID :one
SELECT *
FROM user_addresses
WHERE id = @address_id::uuid LIMIT 1;