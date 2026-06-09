CREATE TABLE user_addresses (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    country_id TEXT NOT NULL,
    country_name TEXT NOT NULL,

    province_id TEXT NOT NULL,
    province_name TEXT NOT NULL,
    
    ward_id TEXT NOT NULL,
    ward_name TEXT NOT NULL,

    address_line TEXT NOT NULL,
    receiver_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    label TEXT NOT NULL DEFAULT 'HOME',
    is_default BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT user_address_label_check
        CHECK (label IN ('OFFICE', 'HOME'))
);

CREATE INDEX idx_user_addresses_user_id ON user_addresses(user_id);
CREATE INDEX idx_user_addresses_country_id ON user_addresses(country_id);
CREATE INDEX idx_user_addresses_province_id ON user_addresses(province_id);
CREATE INDEX idx_user_addresses_ward_id ON user_addresses(ward_id);
