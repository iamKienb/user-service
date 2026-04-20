package otp

import (
	"context"
	"fmt"
	"shopify-user-command-module/internal/application/command/resend_otp"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/application/shared"
	"shopify-user-command-module/internal/domain/identity"
	"time"
)

func (s *otpService) Resend(ctx context.Context, cmd resend_otp.Command) (*resend_otp.Result, error) {
	sessionEntry, err := s.otpCache.GetSession(ctx, cmd.SessionToken)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if sessionEntry == nil {
		return nil, s.wrapError(identity.ErrSessionInvalid)
	}

	count, err := s.otpCache.IncrResendCount(ctx, cmd.SessionToken, shared.ResendWindowTTL)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if count > int64(shared.ResendMaxCount) {
		return nil, s.wrapError(identity.ErrResendLimit)
	}

	otp, err := shared.GenerateOTP()
	if err != nil {
		return nil, s.wrapError(err)
	}

	expiresAt := time.Now().Add(shared.OTPTTL)

	if err := s.otpCache.SaveOTP(ctx, cmd.SessionToken, port.OTPEntry{
		OTP:       otp,
		UserID:    sessionEntry.UserID,
		Email:     sessionEntry.Email,
		ExpiresAt: expiresAt,
		FailCount: 0,
	}, shared.OTPTTL); err != nil {
		return nil, s.wrapError(err)
	}

	fmt.Printf("[OTP] resend email=%s otp=%s expires_in=5min\n", sessionEntry.Email, otp)

	return &resend_otp.Result{
		ExpiresAt: expiresAt,
	}, nil
}
