package otp

import (
	"context"
	"time"

	"shopify-user-command-module/internal/application/command/verify_otp"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/application/shared"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"
)

func (s *otpService) Verify(ctx context.Context, cmd verify_otp.Command) (*verify_otp.Result, error) {
	otpEntry, err := s.validateOTP(ctx, cmd.SessionToken, cmd.OTP)
	if err != nil {
		return nil, s.wrapError(err)
	}

	agg, err := s.activateUserIfNeeded(ctx, otpEntry.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if err := agg.EnsureCredential(); err != nil {
		return nil, s.wrapError(err)
	}

	tokenPair, err := s.tokenGen.GeneratePair(port.TokenClaims{
		UserId:          agg.User.ID.String(),
		Email:           agg.User.Email,
		PasswordVersion: agg.Credential.PasswordVersion,
	})
	if err != nil {
		return nil, s.wrapError(err)
	}

	_ = s.otpCache.DeleteOTP(ctx, cmd.SessionToken)
	_ = s.otpCache.DeleteSession(ctx, cmd.SessionToken)

	return &verify_otp.Result{
		AccessToken:           tokenPair.AccessToken,
		RefreshToken:          tokenPair.RefreshToken,
		AccessTokenExpiresAt:  tokenPair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}, nil
}

func (s *otpService) validateOTP(ctx context.Context, token, inputOTP string) (*port.OTPEntry, error) {
	otpEntry, err := s.otpCache.GetOTP(ctx, token)
	if err != nil {
		return nil, err
	}
	if otpEntry == nil {
		return nil, auth.ErrOTPExpired
	}

	if otpEntry.FailCount >= shared.OTPMaxRetry {
		_ = s.otpCache.DeleteOTP(ctx, token)
		_ = s.otpCache.DeleteSession(ctx, token)
		return nil, auth.ErrOTPMaxRetry
	}

	if inputOTP != otpEntry.OTP {
		otpEntry.FailCount++
		remainingTTL := time.Until(otpEntry.ExpiresAt)
		if remainingTTL > 0 {
			_ = s.otpCache.SaveOTP(ctx, token, *otpEntry, remainingTTL)
		}
		return nil, auth.ErrOTPInvalid
	}

	return otpEntry, nil
}

func (s *otpService) activateUserIfNeeded(ctx context.Context, userID string) (*account.Aggregate, error) {
	agg, err := s.accountRepo.FindAggregateByID(ctx, userID)
	if err != nil || agg == nil {
		return nil, account.ErrUserNotFound
	}

	if agg.ActivateIfNeeded() {
		err := s.txManager.WithTx(ctx, func(txCtx context.Context) error {
			return s.accountRepo.UpdateUser(txCtx, &agg.User)
		})
		if err != nil {
			return nil, err
		}
	}

	return agg, nil
}
