CREATE TABLE shop_member_roles (
    shop_id UUID NOT NULL,
    member_id UUID NOT NULL,
    role_id INT REFERENCES shop_roles(id),
    updated_by UUID REFERENCES users(id) ON DELETE SET NULL,

    PRIMARY KEY (shop_id, member_id, role_id),
    FOREIGN KEY (shop_id, member_id) REFERENCES shop_members(shop_id, member_id) ON DELETE CASCADE
);