-- name: SaveLoginStats :exec
INSERT INTO login_stats (
    user_id,
    failed_count,
    lock_until,
    last_failed_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5)
ON CONFLICT (user_id) DO UPDATE SET
    failed_count = EXCLUDED.failed_count,
    lock_until = EXCLUDED.lock_until,
    last_failed_at = EXCLUDED.last_failed_at,
    updated_at = EXCLUDED.updated_at;

-- name: GetLoginStatsByID :one
SELECT *
FROM login_stats
WHERE user_id = $1;