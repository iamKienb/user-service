package user

import (
	"context"
	"shopify-user-command-module/internal/application/command/login_user"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/identity"
)

func (s *userService) Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error) {
	agg, err := s.repo.FindForLogin(ctx, cmd.Email)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if agg == nil {
		return nil, s.wrapError(identity.ErrUserInvalid)
	}

	stats, err := s.repo.FindLoginStatByID(ctx, cmd.Email)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if stats == nil {
		stats = identity.NewLoginStat(agg.User.ID)
	}

	if stats != nil && !stats.IsLocked() {
		return nil, s.wrapError(identity.ErrAccountLocked)
	}

	check, err := s.hasher.Verify(cmd.Password, agg.Credential.PasswordHash)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if !check {
		stats.RecordFailure()
		if err := s.repo.SaveLoginStat(ctx, stats); err != nil {
			return nil, s.wrapError(err)
		}
		return nil, s.wrapError(identity.ErrInvalidCredentials)
	}

	if !agg.User.IsActive() {
		return nil, s.wrapError(identity.ErrUserNotActive)
	}

	stats.RecordSuccess()
	if err := s.repo.SaveLoginStat(ctx, stats); err != nil {
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

	return &login_user.Result{
		AccessToken:           tokenPair.AccessToken,
		RefreshToken:          tokenPair.RefreshToken,
		AccessTokenExpiresAt:  tokenPair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}, nil

}
