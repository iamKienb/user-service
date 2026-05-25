CREATE TABLE user_addresses (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    
    country_id INT NOT NULL REFERENCES countries(id),
    city_id INT NOT NULL REFERENCES cities(id),
    district_id INT NOT NULL REFERENCES districts(id),
    ward_id INT NOT NULL REFERENCES wards(id),

    address_line TEXT NOT NULL,
    receiver_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    label TEXT NOT NULL DEFAULT 'HOUSE',
    is_default BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    CONSTRAINT user_address_label_check
        CHECK (label IN ('OFFICE', 'HOUSE'))
);

CREATE INDEX idx_user_addresses_user_id ON user_addresses(user_id);
CREATE INDEX idx_user_addresses_country_id ON user_addresses(country_id);
CREATE INDEX idx_user_addresses_city_id ON user_addresses(city_id);
CREATE INDEX idx_user_addresses_district_id ON user_addresses(district_id);
CREATE INDEX idx_user_addresses_ward_id ON user_addresses(ward_id);
