package auth

import (
	"context"
	"errors"
	"fmt"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shared"
	"user-shared-module/common"

	"github.com/jackc/pgx/v5"
)

func (r *authRepository) FindLoginStatByUserID(ctx context.Context, userID shared.UserID) (*auth.LoginStat, error) {
	statsRow, err := r.getQuerier(ctx).GetLoginStatsByID(ctx, common.ToPgUUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get login stats by user id: %w", err)
	}

	return toDomainLoginStat(statsRow), nil
}
