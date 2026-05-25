package login

import (
	"context"
	"fmt"
	"user-command-module/internal/domain/auth"
)

func (r *loginRepository) SaveLoginAttempt(ctx context.Context, stat *auth.LoginAttempt) error {
	if err := r.getQuerier(ctx).SaveLoginStats(ctx, toInfraLoginAttempt(stat)); err != nil {
		return fmt.Errorf("infra: save login stats: %w", err)
	}

	return nil
}
