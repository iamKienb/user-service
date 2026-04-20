package user

import (
	"context"
	"fmt"
	"shopify-user-command-module/internal/application/command/register_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/application/shared"
	"shopify-user-command-module/internal/domain/identity"
	"time"

	"github.com/google/uuid"
)

func (s *userService) Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error) {
	passwordHash, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, s.wrapError(err)
	}

	cacheKey := fmt.Sprintf(shared.EmailCacheKey, cmd.Email)
	existed, err := s.userCache.Exists(ctx, cacheKey)
	if err == nil && existed {
		return nil, s.wrapError(identity.ErrEmailTaken)
	}

	user, err := s.repo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if user != nil {
		_ = s.userCache.Set(ctx, cacheKey, shared.EmailExistsFlag, shared.EmailCacheTTL)
		return nil, s.wrapError(identity.ErrEmailTaken)
	}

	agg := identity.NewAggregate(identity.NewAggregateParams{
		Email:        cmd.Email,
		PasswordHash: passwordHash,
		FullName:     cmd.FullName,
		Gender:       cmd.Gender,
	})

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		return s.repo.Save(ctx, agg)
	}); err != nil {
		return nil, s.wrapError(err)
	}

	s.userCache.Set(ctx, cacheKey, shared.EmailExistsFlag, shared.EmailCacheTTL)

	otp, err := shared.GenerateOTP()
	if err != nil {
		return nil, s.wrapError(err)
	}

	sessionToken := uuid.NewString()
	expiresSessionAt := time.Now().Add(shared.SessionTTL)
	expiresOTPAt := time.Now().Add(shared.OTPTTL)

	if err := s.otpCache.SaveOTP(ctx, sessionToken, port.OTPEntry{
		OTP:       otp,
		UserID:    agg.User.ID.String(),
		Email:     agg.User.Email,
		ExpiresAt: expiresOTPAt,
	}, shared.OTPTTL); err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.otpCache.SaveSession(ctx, sessionToken, port.SessionEntry{
		UserID: agg.User.ID.String(),
		Email:  agg.User.Email,
	}, shared.SessionTTL); err != nil {
		return nil, s.wrapError(err)
	}

	fmt.Printf("[OTP] SEND EMAIL= %s otp= %s\n", agg.User.Email, otp)

	return &register_user.Result{
		SessionToken: sessionToken,
		ExpiresAt:    expiresSessionAt,
	}, nil
}
