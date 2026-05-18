package module

import (
	"context"
	"fmt"

	"user-command-module/internal/application/port"
	"user-command-module/internal/bootstrap/config"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shop"
	"user-command-module/internal/infra/cache"
	accountPg "user-command-module/internal/infra/postgres/account"
	authPg "user-command-module/internal/infra/postgres/auth"
	outboxPg "user-command-module/internal/infra/postgres/outbox"
	shopPg "user-command-module/internal/infra/postgres/shop"
	"user-command-module/internal/infra/security"

	jwtx "github.com/iamKienb/shopify-go-platform/jwt"
	pgx "github.com/iamKienb/shopify-go-platform/postgres"
	redisx "github.com/iamKienb/shopify-go-platform/redis"
)

type InfraModule struct {
	PGService    pgx.PGXService
	RedisService redisx.RedisXService

	AccountRepo account.Repository
	AuthRepo    auth.Repository
	OutboxRepo  port.OutboxRepository
	ShopRepo    shop.Repository

	UserCache port.UserCache
	OtpCache  port.OTPCache
	ShopCache port.ShopCache

	TxManager      port.TxManager
	Hasher         port.PasswordHasher
	TokenGenerator port.TokenService
}

func NewInfraModule(ctx context.Context, cfg *config.UserCommandConfig) (*InfraModule, error) {
	pgService, err := pgx.New(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	redisService, err := redisx.New(cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	jwtService, err := jwtx.New(cfg.Jwt)
	if err != nil {
		return nil, fmt.Errorf("jwt: %w", err)
	}

	return &InfraModule{
		PGService:    pgService,
		RedisService: redisService,

		AccountRepo: accountPg.NewRepository(pgService),
		AuthRepo:    authPg.NewRepository(pgService),
		OutboxRepo:  outboxPg.NewRepository(pgService),
		ShopRepo:    shopPg.NewRepository(pgService),

		UserCache: cache.NewUserCache(redisService),
		OtpCache:  cache.NewOTPCache(redisService),
		ShopCache: cache.NewShopCache(redisService),

		TxManager:      pgx.NewTxManager(pgService.GetPool()),
		Hasher:         security.NewArgon2Hasher(cfg.Argon2),
		TokenGenerator: security.NewTokenGenerator(jwtService),
	}, nil
}
