package user

import (
	"context"
	"shopify-user-command-module/internal/application/command/login_user"
	"shopify-user-command-module/internal/application/command/register_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"
)

type Service interface {
	Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error)
	Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error)
}

type userService struct {
	accountRepo account.Repository
	authRepo    auth.Repository

	userCache port.UserCache
	otpCache  port.OTPCache

	tokenGen  port.TokenService
	txManager port.TxManager
	hasher    port.PasswordHasher
}

func NewUserService(
	accountRepo account.Repository,
	authRepo auth.Repository,
	userCache port.UserCache,
	otpCache port.OTPCache,
	tokenGen port.TokenService,
	txManager port.TxManager,
	hasher port.PasswordHasher,
) Service {
	return &userService{
		accountRepo: accountRepo,
		authRepo:    authRepo,
		userCache:   userCache,
		otpCache:    otpCache,
		tokenGen:    tokenGen,
		txManager:   txManager,
		hasher:      hasher,
	}
}
