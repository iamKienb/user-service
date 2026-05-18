CREATE TABLE user_profiles (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    full_name TEXT NOT NULL,
    gender TEXT NOT NULL DEFAULT 'UNKNOWN',
    phone_number TEXT,
    avatar_url TEXT,
    date_of_birth DATE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ DEFAULT NULL,
    CONSTRAINT user_profile_gender_check
        CHECK (gender IN ('UNKNOWN', 'MALE', 'FEMALE'))
)