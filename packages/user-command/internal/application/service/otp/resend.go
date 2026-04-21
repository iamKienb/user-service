package otp

import (
	"context"
	"fmt"
	"time"

	"shopify-user-command-module/internal/application/command/resend_otp"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/application/shared"
	"shopify-user-command-module/internal/domain/auth"
)

func (s *otpService) Resend(ctx context.Context, cmd resend_otp.Command) (*resend_otp.Result, error) {
	session, err := s.verifyResendPolicy(ctx, cmd.SessionToken)
	if err != nil {
		return nil, s.wrapError(err)
	}

	expiresAt, otp, err := s.generateAndSaveNewOTP(ctx, cmd.SessionToken, session)
	if err != nil {
		return nil, s.wrapError(err)
	}

	fmt.Printf("[RESEND] Email: %s, OTP: %s\n", session.Email, otp)

	return &resend_otp.Result{ExpiresAt: expiresAt}, nil
}

func (s *otpService) verifyResendPolicy(ctx context.Context, token string) (*port.SessionEntry, error) {
	session, err := s.otpCache.GetSession(ctx, token)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, auth.ErrSessionInvalid
	}

	count, err := s.otpCache.IncrResendCount(ctx, token, shared.ResendWindowTTL)
	if err != nil {
		return nil, err
	}
	if count > int64(shared.ResendMaxCount) {
		return nil, auth.ErrResendLimit
	}

	return session, nil
}

func (s *otpService) generateAndSaveNewOTP(ctx context.Context, token string, session *port.SessionEntry) (time.Time, string, error) {
	otp, err := shared.GenerateOTP()
	if err != nil {
		return time.Time{}, "", err
	}

	expiresAt := time.Now().UTC().Add(shared.OTPTTL)
	entry := port.OTPEntry{
		OTP:       otp,
		UserID:    session.UserID,
		Email:     session.Email,
		ExpiresAt: expiresAt,
		FailCount: 0,
	}

	if err := s.otpCache.SaveOTP(ctx, token, entry, shared.OTPTTL); err != nil {
		return time.Time{}, "", err
	}

	return expiresAt, otp, nil
}
