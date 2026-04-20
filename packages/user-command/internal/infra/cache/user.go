package cache

import (
	"context"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/identity"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
	"github.com/redis/go-redis/v9"
)

type userCache struct {
	redisx.Cache
	client *redis.Client
}

func NewUserCache(client *redis.Client) port.UserCache {
	return &userCache{
		Cache:  redisx.NewRedisService(client),
		client: client,
	}
}

func (c *userCache) GetUserInfo(ctx context.Context, key string) (*identity.Profile, error) {
	return nil, nil
}
