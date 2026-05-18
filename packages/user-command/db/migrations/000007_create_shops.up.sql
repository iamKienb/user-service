CREATE TABLE shops (
    id UUID PRIMARY KEY,
    owner_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    slug TEXT NOT NULL UNIQUE,
    status TEXT NOT NULL DEFAULT 'PENDING',

    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    updated_by UUID REFERENCES users(id) ON DELETE SET NULL,
    
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    
    deleted_at TIMESTAMPTZ,
    CONSTRAINT shop_status_check
        CHECK (status IN ('PENDING', 'ACTIVE', 'INACTIVE', 'DELETED'))
);

CREATE TABLE shop_profiles (
    shop_id UUID PRIMARY KEY REFERENCES shops(id) ON DELETE CASCADE,
    description TEXT,
    logo_url TEXT,
    banner_url TEXT,

    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    updated_by UUID REFERENCES users(id) ON DELETE SET NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE INDEX idx_shops_name ON shops(name); 
CREATE INDEX idx_shops_slug ON shops(slug);
CREATE INDEX idx_shops_status ON shops(status);
CREATE INDEX idx_shops_owner_id ON shops(owner_id);