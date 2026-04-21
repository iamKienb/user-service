package port

import (
	"context"
	"time"
)

type UserCache interface {
	IsEmailTaken(ctx context.Context, email string) (bool, error)
	MarkEmailTaken(ctx context.Context, email string, ttl time.Duration) error
}
