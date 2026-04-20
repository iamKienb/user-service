ALTER TABLE outbox
ADD COLUMN idempotency_key UUID,
ADD COLUMN retry_count INT NOT NULL DEFAULT 0,
ADD COLUMN next_retry_at TIMESTAMPTZ,
ADD COLUMN error_message TEXT;

CREATE INDEX idx_outbox_pending ON outbox(status, created_at);
CREATE INDEX idx_outbox_retry ON outbox(next_retry_at) WHERE status = 'FAILED';
CREATE UNIQUE INDEX idx_outbox_idempotency ON outbox(idempotency_key, event_type);
