package login

import (
	"context"
	"errors"
	"fmt"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5"
)

func (r *loginRepository) FindLoginAttemptByID(ctx context.Context, userID shared.UserID) (*auth.LoginAttempt, error) {
	statsRow, err := r.getQuerier(ctx).GetLoginStatsByID(ctx, conv.UUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get login stats by user id: %w", err)
	}

	return toDomainLoginAttempt(statsRow), nil
}
