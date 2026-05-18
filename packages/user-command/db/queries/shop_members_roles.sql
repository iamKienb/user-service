-- name: AddShopMembersBatch :exec
INSERT INTO shop_members (
    shop_id,
    member_id, 
    added_by,
    joined_at
) 
SELECT 
    @shop_id,
    unnest(@member_ids::uuid[]),
    @added_by,
    @joined_at
ON CONFLICT (shop_id, member_id) DO NOTHING;

-- name: AssignShopMemberRolesBatch :exec
INSERT INTO shop_member_roles (
    shop_id,
    member_id,
    role_id,
    updated_by
)
SELECT 
    @shop_id,
    unnest(@member_ids::uuid[]), 
    unnest(@role_ids::int[]),
    @updated_by
ON CONFLICT (shop_id, member_id, role_id) DO UPDATE 
SET updated_by = EXCLUDED.updated_by;

-- name: ClearShopMemberRolesBatch :exec
DELETE FROM shop_member_roles 
WHERE shop_id = $1 
  AND member_id = ANY(@member_ids::uuid[]);

-- name: GetUserRolesInShop :many
SELECT 
    s.id AS shop_id,
    smr.role_id
FROM shops s
LEFT JOIN shop_member_roles smr ON s.id = smr.shop_id AND smr.member_id = $2
WHERE s.id = $1;