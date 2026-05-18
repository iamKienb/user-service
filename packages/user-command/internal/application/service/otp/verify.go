package otp

import (
	"context"
	"time"

	"user-command-module/internal/application/command/verify_otp"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/shared"
	domain_shared "user-command-module/internal/domain/shared"

	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/auth"
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

	tokenPair, err := s.tokenGen.GeneratePair(port.UserClaims{
		UserID:          agg.User.ID.String(),
		Email:           agg.User.Email,
		FullName:        agg.Profile.FullName,
		Roles:           agg.User.Roles,
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
	parseUserID, err := domain_shared.ParseToRawID[domain_shared.UserID](userID)
	if err != nil {
		return nil, account.ErrUserInvalid
	}
	agg, err := s.accountRepo.LoadAggByID(ctx, parseUserID)
	if err != nil || agg == nil {
		return nil, account.ErrUserNotFound
	}

	if agg.ActivateIfVerified() {
		err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
			if err := s.accountRepo.UpdateUser(ctx, &agg.User); err != nil {
				return err
			}

			events := agg.FlushEvents()
			if len(events) == 0 {
				return nil
			}

			publishParams := port.OutboxParam{
				AggregateID:   agg.User.ID.RawID(),
				AggregateType: account.AggregateTypeUser,
				Events:        events,
			}

			return s.outboxService.Publish(ctx, publishParams)
		})
		if err != nil {
			return nil, s.wrapError(err)
		}
	}

	return agg, nil
}
