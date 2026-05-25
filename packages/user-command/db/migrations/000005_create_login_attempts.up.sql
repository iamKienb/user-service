CREATE TABLE login_attempts (
    user_id UUID PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    failed_count INT NOT NULL DEFAULT 0,
    lock_until TIMESTAMPTZ,
    last_failed_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)