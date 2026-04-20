CREATE TABLE shop_members (
    shop_id UUID REFERENCES shops(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id) ON DELETE CASCADE, 
    role_id INT NOT NULL, 
    PRIMARY KEY (shop_id, user_id)
);
