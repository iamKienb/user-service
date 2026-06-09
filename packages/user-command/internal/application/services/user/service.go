package user

import (
	"context"
	"user-command-module/internal/application/commands/add_user_address"
	"user-command-module/internal/application/commands/login_user"
	"user-command-module/internal/application/commands/register_user"
	"user-command-module/internal/application/port"
	get_user_address_by_id "user-command-module/internal/application/queries/get_address_by_id"
	"user-command-module/internal/application/services/outbox"
	"user-command-module/internal/domain/address"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/profile"
	"user-command-module/internal/domain/shared"
	domain_user "user-command-module/internal/domain/user"
)

type Service interface {
	Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error)
	Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error)
	AddAddress(ctx context.Context, cmd add_user_address.Command) (*add_user_address.Result, error)
	GetAddress(ctx context.Context, qry get_user_address_by_id.Query) (*get_user_address_by_id.Result, error)
}

type userService struct {
	userRepo        domain_user.Repository
	authRepo        auth.Repository
	profileRepo     profile.Repository
	userAddressRepo address.Repository

	outboxService outbox.Service

	userCache port.UserCache
	otpCache  port.OTPCache

	txManager port.TxManager
	tokenGen  port.TokenService
	hasher    port.PasswordHasher
}

func NewUserService(
	userRepo domain_user.Repository,
	authRepo auth.Repository,
	profileRepo profile.Repository,
	userAddressRepo address.Repository,

	outboxService outbox.Service,

	userCache port.UserCache,
	otpCache port.OTPCache,

	tokenGen port.TokenService,
	txManager port.TxManager,
	hasher port.PasswordHasher,
) Service {
	return &userService{
		userRepo:        userRepo,
		authRepo:        authRepo,
		profileRepo:     profileRepo,
		userAddressRepo: userAddressRepo,

		outboxService: outboxService,

		userCache: userCache,
		otpCache:  otpCache,

		tokenGen:  tokenGen,
		txManager: txManager,
		hasher:    hasher,
	}
}

func (s *userService) validateAndCheckActiveUser(ctx context.Context, userID shared.UserID) (*domain_user.User, error) {
	user, err := s.userRepo.FindUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain_user.ErrUserNotFound
	}
	if err := user.EnsureActive(); err != nil {
		return nil, err
	}
	return user, nil
}
