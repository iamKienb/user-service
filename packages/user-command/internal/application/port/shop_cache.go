package port

import (
	"context"
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopCache interface {
	GetSlugFromBloomFilter(ctx context.Context, slug string) (int, error)
	AddSlugToBloomFilter(ctx context.Context, slug string) error
	IsIdemKeyTaken(ctx context.Context, userID shared.UserID) (bool, error)
	SetIdemKey(ctx context.Context, userID shared.UserID, ttl time.Duration) error
}
