package module

import (
	"context"
	"fmt"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/bootstrap/config"
	"shopify-user-command-module/internal/domain/identity"
	"shopify-user-command-module/internal/infra/cache"
	"shopify-user-command-module/internal/infra/postgres/user"
	"shopify-user-command-module/internal/infra/security"

	postgresx "github.com/iamKienb/shopify-go-platform/postgres"
	redisx "github.com/iamKienb/shopify-go-platform/redis"
	"github.com/jackc/pgx/v5/pgxpool"
)

type InfraModule struct {
	Pool           *pgxpool.Pool
	IdentityRepo   identity.Repository
	Cache          port.UserCache
	OtpCache       port.OTPCache
	TxManager      port.TxManager
	Hasher         port.PasswordHasher
	TokenGenerator port.TokenGenerator
}

func NewInfraModule(ctx context.Context, cfg *config.UserCommandConfig) (*InfraModule, error) {
	pgClient, err := postgresx.New(cfg.Postgres)
	if err != nil {
		return nil, fmt.Errorf("Postgres: %w", err)
	}

	redisClient, err := redisx.New(cfg.Redis)
	if err != nil {
		return nil, fmt.Errorf("Redis: %w", err)
	}

	repo := user.NewUserRepository(pgClient.Pool)

	userCache := cache.NewUserCache(redisClient.Conn)
	otpCache := cache.NewOTPCache(redisClient.Conn)

	txManager := postgresx.NewTxManager(pgClient.Pool)

	agronHasher := security.NewArgon2Hasher(cfg.Argon2)

	jwtGenerator := security.NewJWTGenerator(cfg.Jwt)

	return &InfraModule{
		Pool:           pgClient.Pool,
		IdentityRepo:   repo,
		Cache:          userCache,
		OtpCache:       otpCache,
		TxManager:      txManager,
		Hasher:         agronHasher,
		TokenGenerator: jwtGenerator,
	}, nil

}
