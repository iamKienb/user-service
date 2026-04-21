package user

import (
	"context"
	"shopify-user-command-module/internal/application/command/login_user"
	"shopify-user-command-module/internal/application/command/register_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/identity"
)

type Service interface {
	Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error)
	Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error)
}

type userService struct {
	repo identity.Repository

	userCache port.UserCache
	otpCache  port.OTPCache

	tokenGen  port.TokenGenerator
	txManager port.TxManager
	hasher    port.PasswordHasher
}

func NewUserService(
	repo identity.Repository,
	userCache port.UserCache,
	otpCache port.OTPCache,
	tokenGen port.TokenGenerator,
	txManager port.TxManager,
	hasher port.PasswordHasher,
) Service {
	return &userService{
		repo:      repo,
		userCache: userCache,
		otpCache:  otpCache,
		tokenGen:  tokenGen,
		txManager: txManager,
		hasher:    hasher,
	}
}
