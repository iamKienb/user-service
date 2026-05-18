package account

import (
	"context"
	"errors"
	"fmt"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/shared"
	"user-shared-module/common"

	"github.com/jackc/pgx/v5"
)

func (r *accountRepository) LoadAggByID(ctx context.Context, userID shared.UserID) (*account.Aggregate, error) {
	pgUUID := common.ToPgUUID(userID)
	q := r.getQuerier(ctx)

	userRow, err := q.GetUserByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by id: %w", err)
	}

	credentialRow, err := q.GetUserCredentialByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get credential by id: %w", err)
	}

	profileRow, err := q.GetUserProfileByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get profile by id: %w", err)
	}

	return account.LoadAggregate(
		*toDomainUser(userRow),
		toDomainCredential(credentialRow),
		toDomainProfile(profileRow),
	), nil
}

func (r *accountRepository) FindByUserID(ctx context.Context, userID shared.UserID) (*account.User, error) {
	userRow, err := r.getQuerier(ctx).GetUserByID(ctx, common.ToPgUUID(userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by id: %w", err)
	}

	return toDomainUser(userRow), nil
}

func (r *accountRepository) FindByEmail(ctx context.Context, email string) (*account.User, error) {
	userRow, err := r.getQuerier(ctx).GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by email: %w", err)
	}

	return toDomainUser(userRow), nil
}

func (r *accountRepository) LoadAggByEmail(ctx context.Context, email string) (*account.Aggregate, error) {
	q := r.getQuerier(ctx)

	userRow, err := q.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user for login: %w", err)
	}

	credentialRow, err := q.GetUserCredentialByID(ctx, userRow.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get credential for login: %w", err)
	}

	return account.LoadAggregate(
		*toDomainUser(userRow),
		toDomainCredential(credentialRow),
		nil,
	), nil
}
