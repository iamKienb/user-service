CREATE TABLE addresses (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    label TEXT NOT NULL,
    receiver_name TEXT NOT NULL,
    phone_number TEXT NOT NULL,
    address_line TEXT NOT NULL,
    ward TEXT NOT NULL,
    district TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    is_default BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_At TIMESTAMPTZ NOT NULL DEFAULT now()
)
