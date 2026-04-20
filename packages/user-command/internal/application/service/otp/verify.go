package otp

import (
	"context"
	"shopify-user-command-module/internal/application/command/verify_otp"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/application/shared"
	"shopify-user-command-module/internal/domain/identity"
	"time"
)

func (s *otpService) Verify(ctx context.Context, cmd verify_otp.Command) (*verify_otp.Result, error) {
	sessionEntry, err := s.otpCache.GetSession(ctx, cmd.SessionToken)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if sessionEntry == nil {
		return nil, s.wrapError(identity.ErrSessionInvalid)
	}

	otpEntry, err := s.otpCache.GetOTP(ctx, cmd.SessionToken)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if otpEntry == nil {
		return nil, s.wrapError(identity.ErrOTPExpired)
	}

	if otpEntry.FailCount >= shared.OTPMaxRetry {
		_ = s.otpCache.DeleteOTP(ctx, cmd.SessionToken)
		_ = s.otpCache.DeleteSession(ctx, cmd.SessionToken)
		return nil, s.wrapError(identity.ErrOTPMaxRetry)
	}

	remainingTTL := time.Until(otpEntry.ExpiresAt)
	if remainingTTL <= 0 {
		return nil, s.wrapError(identity.ErrOTPExpired)
	}

	if cmd.OTP != otpEntry.OTP {
		otpEntry.FailCount++
		_ = s.otpCache.SaveOTP(ctx, cmd.SessionToken, *otpEntry, remainingTTL)
		return nil, s.wrapError(identity.ErrOTPInvalid)
	}

	var aggActive *identity.IdentityAggregate
	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		agg, err := s.repo.FindAggregateByID(ctx, otpEntry.UserID)
		if err != nil {
			return err
		}

		if !agg.User.IsActive() {
			agg.User.Activate()
		}

		if err := s.repo.UpdateUser(ctx, &agg.User); err != nil {
			return err
		}
		aggActive = agg
		return nil
	}); err != nil {
		return nil, s.wrapError(err)
	}

	tokenPair, err := s.tokenGen.GeneratePair(port.TokenClaims{
		UserId:          aggActive.User.ID.String(),
		Email:           aggActive.User.Email,
		PasswordVersion: aggActive.Credential.PasswordVersion,
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
