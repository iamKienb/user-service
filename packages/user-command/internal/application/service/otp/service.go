package otp

import (
	"context"
	"shopify-user-command-module/internal/application/command/resend_otp"
	"shopify-user-command-module/internal/application/command/verify_otp"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/account"
)

type Service interface {
	Verify(ctx context.Context, cmd verify_otp.Command) (*verify_otp.Result, error)
	Resend(ctx context.Context, cmd resend_otp.Command) (*resend_otp.Result, error)
}

type otpService struct {
	accountRepo account.Repository
	tokenGen    port.TokenGenerator
	otpCache    port.OTPCache
	txManager   port.TxManager
}

func NewOTPService(accountRepo account.Repository, tokenGen port.TokenGenerator, otpCache port.OTPCache, txManager port.TxManager) Service {
	return &otpService{
		accountRepo: accountRepo,
		tokenGen:    tokenGen,
		otpCache:    otpCache,
		txManager:   txManager,
	}
}
