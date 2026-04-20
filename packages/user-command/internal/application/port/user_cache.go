package port

import (
	"context"
	"shopify-user-command-module/internal/domain/identity"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
)

type UserCache interface {
	redisx.Cache
	GetUserInfo(ctx context.Context, key string) (*identity.Profile, error)
}
