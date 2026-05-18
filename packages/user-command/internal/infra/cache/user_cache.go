package cache

import (
	"context"
	"fmt"
	"time"
	"user-command-module/internal/application/port"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
)

const emailCacheKey = "user-command:user:email:%s"

type userCache struct {
	cache redisx.Cache
}

func NewUserCache(service redisx.RedisXService) port.UserCache {
	return &userCache{
		cache: service,
	}
}

func (c *userCache) IsEmailTaken(ctx context.Context, email string) (bool, error) {
	return c.cache.Exists(ctx, fmt.Sprintf(emailCacheKey, email))
}

func (c *userCache) MarkEmailTaken(ctx context.Context, email string, ttl time.Duration) error {
	return c.cache.Set(ctx, fmt.Sprintf(emailCacheKey, email), "exists", ttl)
}
