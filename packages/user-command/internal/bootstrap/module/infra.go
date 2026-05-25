package module

import (
	"context"
	"fmt"

	"user-command-module/internal/application/port"
	"user-command-module/internal/bootstrap/config"
	"user-command-module/internal/domain/address"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/profile"
	"user-command-module/internal/domain/user"
	"user-command-module/internal/infra/cache"
	addressPg "user-command-module/internal/infra/postgres/address"
	loginPg "user-command-module/internal/infra/postgres/login"
	profilePg "user-command-module/internal/infra/postgres/profile"
	userPg "user-command-module/internal/infra/postgres/user"

	outboxPg "user-command-module/internal/infra/postgres/outbox"
	"user-command-module/internal/infra/security"

	jwtx "github.com/iamKienb/go-core/jwt"
	pgx "github.com/iamKienb/go-core/postgres"
	redisx "github.com/iamKienb/go-core/redis"
)

type InfraModule struct {
	PGService    pgx.PGXService
	RedisService redisx.RedisXService

	UserRepo        user.Repository
	AuthRepo        auth.Repository
	ProfileRepo     profile.Repository
	UserAddressRepo address.Repository
	OutboxRepo      port.OutboxRepository

	UserCache port.UserCache
	OtpCache  port.OTPCache

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

		UserRepo:        userPg.NewRepository(pgService),
		AuthRepo:        loginPg.NewRepository(pgService),
		ProfileRepo:     profilePg.NewRepository(pgService),
		UserAddressRepo: addressPg.NewRepository(pgService),
		OutboxRepo:      outboxPg.NewRepository(pgService),

		UserCache: cache.NewUserCache(redisService),
		OtpCache:  cache.NewOTPCache(redisService),

		TxManager:      pgx.NewTxManager(pgService.GetPool()),
		Hasher:         security.NewArgon2Hasher(cfg.Argon2),
		TokenGenerator: security.NewTokenGenerator(jwtService),
	}, nil
}
