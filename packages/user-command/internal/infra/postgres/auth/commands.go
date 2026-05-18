package auth

import (
	"context"
	"fmt"
	"user-command-module/internal/domain/auth"
)

func (r *authRepository) SaveLoginStat(ctx context.Context, stat *auth.LoginStat) error {
	if err := r.getQuerier(ctx).SaveLoginStats(ctx, toInfraLoginStat(stat)); err != nil {
		return fmt.Errorf("infra: save login stats: %w", err)
	}

	return nil
}
