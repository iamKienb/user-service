package otp

import (
	"context"
	"user-command-module/internal/application/commands/resend_otp"
	"user-command-module/internal/application/commands/verify_otp"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/services/outbox"
	"user-command-module/internal/domain/profile"
	"user-command-module/internal/domain/user"
)

type Service interface {
	Verify(ctx context.Context, cmd verify_otp.Command) (*verify_otp.Result, error)
	Resend(ctx context.Context, cmd resend_otp.Command) (*resend_otp.Result, error)
}

type otpService struct {
	userRepo      user.Repository
	profileRepo   profile.Repository
	outboxService outbox.Service

	tokenGen  port.TokenService
	otpCache  port.OTPCache
	txManager port.TxManager
}

func NewOTPService(
	userRepo user.Repository,
	profileRepo profile.Repository,
	outboxService outbox.Service,

	tokenGen port.TokenService,
	otpCache port.OTPCache,
	txManager port.TxManager,
) Service {
	return &otpService{
		userRepo:      userRepo,
		profileRepo:   profileRepo,
		outboxService: outboxService,

		tokenGen:  tokenGen,
		otpCache:  otpCache,
		txManager: txManager,
	}
}
