package user

import (
	"context"
	"errors"
	"fmt"
	"user-command-module/internal/domain/shared"
	domain_user "user-command-module/internal/domain/user"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5"
)

func (r *userRepository) FindUserByID(ctx context.Context, userID shared.UserID) (*domain_user.User, error) {
	userRow, err := r.getQuerier(ctx).FindUserByID(ctx, conv.UUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by id: %w", err)
	}

	// credentialRow, err := r.getQuerier(ctx).FindUserCredentialByID(ctx, userRow.ID)
	// if err != nil {
	// 	if errors.Is(err, pgx.ErrNoRows) {
	// 		return nil, nil
	// 	}
	// 	return nil, fmt.Errorf("infra: get credential by id: %w", err)
	// }

	return toDomainUser(userRow, nil), nil
}

func (r *userRepository) FindUserByEmail(ctx context.Context, email string) (*domain_user.User, error) {
	userRow, err := r.getQuerier(ctx).FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by email: %w", err)
	}

	credentialRow, err := r.getQuerier(ctx).FindUserCredentialByID(ctx, userRow.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get credential by id: %w", err)
	}

	return toDomainUser(userRow, &credentialRow), nil
}
