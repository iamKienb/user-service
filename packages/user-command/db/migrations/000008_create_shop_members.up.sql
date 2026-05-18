CREATE TABLE shop_members (
    shop_id UUID REFERENCES shops(id) ON DELETE CASCADE,
    member_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    joined_at TIMESTAMPTZ DEFAULT now(),
    added_by  UUID REFERENCES users(id) ON DELETE SET NULL,
    PRIMARY KEY (shop_id, member_id)
);
