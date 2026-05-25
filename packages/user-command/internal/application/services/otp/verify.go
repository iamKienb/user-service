package otp

import (
	"context"
	"time"

	"user-command-module/internal/application/commands/verify_otp"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/shared"
	"user-command-module/internal/domain/auth"
	domain_shared "user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/user"
)

func (s *otpService) Verify(ctx context.Context, cmd verify_otp.Command) (*verify_otp.Result, error) {
	otpEntry, err := s.validateOTP(ctx, cmd.SessionToken, cmd.OTP)
	if err != nil {
		return nil, s.wrapError(err)
	}

	user, err := s.activateUserIfNeeded(ctx, otpEntry.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	profile, err := s.profileRepo.FindProfileByID(ctx, user.ID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	tokenPair, err := s.tokenGen.GeneratePair(port.UserClaims{
		UserID:          user.ID.String(),
		Email:           user.Email,
		FullName:        profile.FullName,
		Roles:           user.Roles,
		PasswordVersion: user.Credential.PasswordVersion,
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

func (s *otpService) activateUserIfNeeded(ctx context.Context, userID string) (*user.User, error) {
	parseUserID, err := domain_shared.ParseToRawID[domain_shared.UserID](userID)
	if err != nil {
		return nil, user.ErrUserInvalid
	}

	u, err := s.userRepo.FindUserByID(ctx, parseUserID)
	if err != nil || u == nil {
		return nil, user.ErrUserNotFound
	}

	if u.ActivateIfVerified() {
		err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
			if err := s.userRepo.UpdateUser(ctx, u); err != nil {
				return err
			}

			if events := u.FlushEvents(); len(events) > 0 {
				if err := s.outboxService.Publish(ctx, port.OutboxParam{
					AggregateID:   u.ID.RawID(),
					AggregateType: u.Type(),
					Events:        events,
				}); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return nil, s.wrapError(err)
		}
	}

	return u, nil
}
