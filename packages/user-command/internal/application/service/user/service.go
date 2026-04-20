package user

import (
	"context"
	"shopify-user-command-module/internal/application/command/register_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/identity"
)

type Service interface {
	Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error)
}

type userService struct {
	repo      identity.Repository
	userCache port.UserCache
	otpCache  port.OTPCache
	txManager port.TxManager
	hasher    port.PasswordHasher
}

func NewUserService(
	repo identity.Repository,
	userCache port.UserCache,
	otpCache port.OTPCache,
	txManager port.TxManager,
	hasher port.PasswordHasher,
) Service {
	return &userService{
		repo:      repo,
		userCache: userCache,
		otpCache:  otpCache,
		txManager: txManager,
		hasher:    hasher,
	}
}
