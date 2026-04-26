package module

import (
	"context"
	"fmt"

	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/bootstrap/config"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"
	"shopify-user-command-module/internal/infra/cache"
	accountpg "shopify-user-command-module/internal/infra/postgres/account"
	authpg "shopify-user-command-module/internal/infra/postgres/auth"
	"shopify-user-command-module/internal/infra/security"

	postgresx "github.com/iamKienb/shopify-go-platform/postgres"
	redisx "github.com/iamKienb/shopify-go-platform/redis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfraModule struct {
	PostgresPool   *pgxpool.Pool
	AccountRepo    account.Repository
	AuthRepo       auth.Repository
	UserCache      port.UserCache
	OtpCache       port.OTPCache
	TxManager      port.TxManager
	Hasher         port.PasswordHasher
	TokenGenerator port.TokenService
}

func NewInfraModule(ctx context.Context, cfg *config.UserCommandConfig) (*InfraModule, error) {
	pgClient, err := postgresx.New(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	redisClient, err := redisx.New(cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	return &InfraModule{
		PostgresPool:   pgClient.Pool,
		AccountRepo:    accountpg.NewRepository(pgClient.Pool),
		AuthRepo:       authpg.NewRepository(pgClient.Pool),
		UserCache:      cache.NewUserCache(redisClient.Conn),
		OtpCache:       cache.NewOTPCache(redisClient.Conn),
		TxManager:      postgresx.NewTxManager(pgClient.Pool),
		Hasher:         security.NewArgon2Hasher(cfg.Argon2),
		TokenGenerator: security.NewTokenGenerator(cfg.Jwt),
	}, nil
}
