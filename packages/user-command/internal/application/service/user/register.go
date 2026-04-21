package user

import (
	"context"
	"fmt"
	"time"

	"shopify-user-command-module/internal/application/command/register_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/application/shared"
	"shopify-user-command-module/internal/domain/account"

	"github.com/google/uuid"
)

func (s *userService) Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error) {
	if err := s.ensureEmailAvailable(ctx, cmd.Email); err != nil {
		return nil, s.wrapError(err)
	}

	passwordHash, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, s.wrapError(err)
	}

	validGender, err := account.ValidateGender(cmd.Gender)
	if err != nil {
		return nil, s.wrapError(err)
	}

	agg := account.NewAggregate(account.NewAggregateParams{
		Email:        cmd.Email,
		PasswordHash: passwordHash,
		FullName:     cmd.FullName,
		Gender:       validGender,
	})

	if err := s.txManager.WithTx(ctx, func(txCtx context.Context) error {
		return s.accountRepo.SaveAggregate(txCtx, agg)
	}); err != nil {
		return nil, s.wrapError(err)
	}
	_ = s.userCache.MarkEmailTaken(ctx, cmd.Email, shared.EmailCacheTTL)

	res, err := s.createRegistrationChallenge(ctx, agg)
	if err != nil {
		return nil, s.wrapError(err)
	}

	return res, nil
}

func (s *userService) ensureEmailAvailable(ctx context.Context, email string) error {
	taken, _ := s.userCache.IsEmailTaken(ctx, email)
	if taken {
		return account.ErrEmailTaken
	}

	existing, err := s.accountRepo.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existing != nil {
		_ = s.userCache.MarkEmailTaken(ctx, email, shared.EmailCacheTTL)
		return account.ErrEmailTaken
	}
	return nil
}

func (s *userService) createRegistrationChallenge(ctx context.Context, agg *account.Aggregate) (*register_user.Result, error) {
	otp, err := shared.GenerateOTP()
	if err != nil {
		return nil, err
	}

	sessionToken := uuid.NewString()
	now := time.Now().UTC()

	if err := s.otpCache.SaveOTP(ctx, sessionToken, port.OTPEntry{
		OTP:       otp,
		UserID:    agg.User.ID.String(),
		Email:     agg.User.Email,
		ExpiresAt: now.Add(shared.OTPTTL),
	}, shared.OTPTTL); err != nil {
		return nil, err
	}

	if err := s.otpCache.SaveSession(ctx, sessionToken, port.SessionEntry{
		UserID: agg.User.ID.String(),
		Email:  agg.User.Email,
	}, shared.SessionTTL); err != nil {
		return nil, err
	}

	fmt.Printf("[DEBUG] SESSION: %s, OTP: %s\n", sessionToken, otp)

	return &register_user.Result{
		SessionToken: sessionToken,
		ExpiresAt:    now.Add(shared.SessionTTL),
	}, nil
}
