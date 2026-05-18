package user

import (
	"context"
	"user-command-module/internal/application/command/add_user_address"
	"user-command-module/internal/application/command/login_user"
	"user-command-module/internal/application/command/register_user"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/service/outbox"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/auth"
)

type Service interface {
	Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error)
	Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error)
	AddAddress(ctx context.Context, cmd add_user_address.Command) (*add_user_address.Result, error)
}

type userService struct {
	accountRepo   account.Repository
	authRepo      auth.Repository
	outboxService outbox.Service

	userCache port.UserCache
	otpCache  port.OTPCache

	txManager port.TxManager
	tokenGen  port.TokenService
	hasher    port.PasswordHasher
}

func NewUserService(
	accountRepo account.Repository,
	authRepo auth.Repository,
	outboxService outbox.Service,

	userCache port.UserCache,
	otpCache port.OTPCache,

	tokenGen port.TokenService,
	txManager port.TxManager,
	hasher port.PasswordHasher,
) Service {
	return &userService{
		accountRepo:   accountRepo,
		authRepo:      authRepo,
		outboxService: outboxService,

		userCache: userCache,
		otpCache:  otpCache,

		tokenGen:  tokenGen,
		txManager: txManager,
		hasher:    hasher,
	}
}
