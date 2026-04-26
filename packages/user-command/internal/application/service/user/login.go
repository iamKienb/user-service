package user

import (
	"context"

	"shopify-user-command-module/internal/application/command/login_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"
)

func (s *userService) Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error) {
	agg, err := s.accountRepo.FindForAuthentication(ctx, cmd.Email)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if agg == nil {
		return nil, s.wrapError(auth.ErrInvalidCredentials)
	}

	if err := s.verifyLoginPolicy(ctx, agg, cmd.Password); err != nil {
		return nil, s.wrapError(err)
	}

	tokenPair, err := s.tokenGen.GeneratePair(port.UserClaims{
		UserID:          agg.User.ID.String(),
		Email:           agg.User.Email,
		Roles:           agg.User.Roles,
		PasswordVersion: agg.Credential.PasswordVersion,
	})
	if err != nil {
		return nil, s.wrapError(err)
	}

	return &login_user.Result{
		AccessToken:           tokenPair.AccessToken,
		RefreshToken:          tokenPair.RefreshToken,
		AccessTokenExpiresAt:  tokenPair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}, nil
}

func (s *userService) verifyLoginPolicy(ctx context.Context, agg *account.Aggregate, password string) error {
	if err := agg.EnsureCredential(); err != nil {
		return err
	}

	stats, err := s.authRepo.FindLoginStatByUserID(ctx, agg.User.ID.String())
	if err != nil {
		return err
	}
	if stats == nil {
		stats = auth.NewLoginStat(agg.User.ID)
	}

	if err := stats.EnsureCanAttemptLogin(); err != nil {
		return err
	}

	matched, err := s.hasher.Verify(password, agg.Credential.PasswordHash)
	if err != nil || !matched {
		stats.RecordFailure()
		_ = s.authRepo.SaveLoginStat(ctx, stats)
		return auth.ErrInvalidCredentials
	}

	if err := agg.EnsureCanLogin(); err != nil {
		return err
	}

	stats.RecordSuccess()

	return s.authRepo.SaveLoginStat(ctx, stats)
}
