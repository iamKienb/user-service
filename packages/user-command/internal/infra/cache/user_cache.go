package cache

import (
	"context"
	"fmt"
	"time"

	"shopify-user-command-module/internal/application/port"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
	"github.com/redis/go-redis/v9"
)

const emailCacheKey = "user-command:email:%s"

type userCache struct {
	cache redisx.Cache
}

func NewUserCache(client *redis.Client) port.UserCache {
	return &userCache{
		cache: redisx.NewRedisService(client),
	}
}

func (c *userCache) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	return c.cache.Exists(ctx, fmt.Sprintf(emailCacheKey, email))
}

func (c *userCache) MarkEmailTaken(ctx context.Context, email string, ttl time.Duration) error {
	return c.cache.Set(ctx, fmt.Sprintf(emailCacheKey, email), "exists", ttl)
}
