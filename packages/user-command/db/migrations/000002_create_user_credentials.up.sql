CREATE TABLE user_credentials (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    password_hash TEXT NOT NULL,
    password_version INT NOT NULL DEFAULT 1,
    password_updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)
