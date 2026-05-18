CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    email_verified_at TIMESTAMPTZ,
    status TEXT NOT NULL DEFAULT 'PENDING',
    roles TEXT[] NOT NULL DEFAULT '{CUSTOMER}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    deleted_at TIMESTAMPTZ DEFAULT NULL,
    CONSTRAINT user_status_check
        CHECK (status IN ('PENDING', 'ACTIVE', 'BANNED', 'DELETED')),
    CONSTRAINT user_roles_check
        CHECK (roles <@ ARRAY['SUPER_ADMIN', 'SHOP_OWNER', 'SHOP_STAFF', 'CUSTOMER']::TEXT[])
)