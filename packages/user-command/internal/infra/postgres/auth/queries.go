package auth

import (
	"context"
	"errors"
	"fmt"

	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"
	"shopify-user-command-module/internal/infra/common"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *authRepository) FindLoginStatByUserID(ctx context.Context, userID string) (*auth.LoginStat, error) {
	parsedID, err := uuid.Parse(userID)
	if err != nil {
		return nil, account.ErrUserInvalid
	}

	statsRow, err := r.getQuerier(ctx).GetLoginStatsByID(ctx, common.ToPgUUID(parsedID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get login stats by user id: %w", err)
	}

	return toDomainLoginStat(statsRow), nil
}
