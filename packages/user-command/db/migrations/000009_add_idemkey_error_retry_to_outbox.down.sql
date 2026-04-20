ALTER TABLE outbox
DROP COLUMN IF EXISTS idempotency_key,
DROP COLUMN IF EXISTS retry_count,
DROP COLUMN IF EXISTS next_retry_at,
DROP COLUMN IF EXISTS error_message;
