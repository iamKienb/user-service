package profile

import (
	"context"
	"errors"
	"fmt"
	"user-command-module/internal/domain/profile"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5"
)

func (r *profileRepository) FindProfileByID(ctx context.Context, userID shared.UserID) (*profile.Profile, error) {
	profileRow, err := r.getQuerier(ctx).FindUserProfileByID(ctx, conv.UUID(userID))
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("infra: get profile for login: %w", err)
	}

	return toDomainProfile(profileRow), nil
}
