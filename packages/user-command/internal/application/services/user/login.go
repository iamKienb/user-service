package user

import (
	"context"

	"user-command-module/internal/application/commands/login_user"
	"user-command-module/internal/application/port"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/user"
)

func (s *userService) Login(ctx context.Context, cmd login_user.Command) (*login_user.Result, error) {
	user, err := s.userRepo.FindUserByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.verifyLoginPolicy(ctx, user, cmd.Password); err != nil {
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

	return &login_user.Result{
		AccessToken:           tokenPair.AccessToken,
		RefreshToken:          tokenPair.RefreshToken,
		AccessTokenExpiresAt:  tokenPair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: tokenPair.RefreshTokenExpiresAt,
	}, nil
}

func (s *userService) verifyLoginPolicy(ctx context.Context, user *user.User, password string) error {
	attempt, err := s.authRepo.FindLoginAttemptByID(ctx, user.ID)
	if err != nil {
		return err
	}
	if attempt == nil {
		attempt = auth.NewAttempt(user.ID)
	}

	if err := attempt.EnsureCanAttemptLogin(); err != nil {
		return err
	}

	matched, err := s.hasher.Verify(password, user.Credential.PasswordHash)
	if err != nil || !matched {
		attempt.RecordFailure()
		_ = s.authRepo.SaveLoginAttempt(ctx, attempt)
		return auth.ErrInvalidCredentials
	}

	if err := user.EnsureActiveForLogin(); err != nil {
		return err
	}

	attempt.RecordSuccess()

	return s.authRepo.SaveLoginAttempt(ctx, attempt)
}
