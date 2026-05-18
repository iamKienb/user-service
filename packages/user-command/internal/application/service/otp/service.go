package otp

import (
	"context"
	"user-command-module/internal/application/command/resend_otp"
	"user-command-module/internal/application/command/verify_otp"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/service/outbox"
	"user-command-module/internal/domain/account"
)

type Service interface {
	Verify(ctx context.Context, cmd verify_otp.Command) (*verify_otp.Result, error)
	Resend(ctx context.Context, cmd resend_otp.Command) (*resend_otp.Result, error)
}

type otpService struct {
	accountRepo   account.Repository
	outboxService outbox.Service

	tokenGen  port.TokenService
	otpCache  port.OTPCache
	txManager port.TxManager
}

func NewOTPService(
	accountRepo account.Repository,
	outboxService outbox.Service,

	tokenGen port.TokenService,
	otpCache port.OTPCache,
	txManager port.TxManager,
) Service {
	return &otpService{
		accountRepo:   accountRepo,
		outboxService: outboxService,

		tokenGen:  tokenGen,
		otpCache:  otpCache,
		txManager: txManager,
	}
}
