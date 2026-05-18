package cache

import (
	"context"
	"fmt"
	"time"
	"user-command-module/internal/application/port"
	"user-command-module/internal/domain/shared"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
	"github.com/redis/go-redis/v9"
)

const filter string = "shops_slug"
const idemKey = "user-command:shop:key:%s"

type shopCache struct {
	cache  redisx.RedisXService
	client *redis.Client
}

func NewShopCache(service redisx.RedisXService) port.ShopCache {
	return &shopCache{
		cache:  service,
		client: service.GetClient(),
	}
}

func (c *shopCache) IsIdemKeyTaken(ctx context.Context, userID shared.UserID) (bool, error) {
	return c.cache.Exists(ctx, fmt.Sprintf(idemKey, userID))
}

func (c *shopCache) SetIdemKey(ctx context.Context, userID shared.UserID, ttl time.Duration) error {
	return c.cache.Set(ctx, fmt.Sprintf(idemKey, userID), "exists", ttl)
}

func (c *shopCache) AddSlugToBloomFilter(ctx context.Context, slug string) error {
	if err := c.InitBloomFilter(ctx, c.client); err != nil {
		return fmt.Errorf("failed to init bloom filter: %w", err)
	}

	err := c.client.Do(ctx, "BF.ADD", filter, slug).Err()
	if err != nil {
		return fmt.Errorf("redis error adding slug to bloom filter: %w", err)
	}

	return nil
}

func (c *shopCache) GetSlugFromBloomFilter(ctx context.Context, slug string) (int, error) {
	exists, err := c.client.Do(ctx, "BF.EXISTS", filter, slug).Int()
	if err != nil {
		return 0, fmt.Errorf("redis error: %w", err)
	}

	return exists, nil
}

func (c *shopCache) InitBloomFilter(ctx context.Context, client *redis.Client) error {
	err := client.Do(ctx, "BF.RESERVE", filter, "0.01", "1000000").Err()
	if err != nil && err.Error() != "ERR item already exists" {
		return err
	}

	return nil
}
