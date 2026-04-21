package user

import (
	"context"
	"errors"
	"fmt"
	"shopify-user-command-module/internal/domain/identity"
	"shopify-user-command-module/internal/infra/common"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (r *userRepository) FindAggregateByID(ctx context.Context, id string) (*identity.IdentityAggregate, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, identity.ErrUserInvalid
	}

	pgUUID := common.ToPgUUID(uuid)

	user, err := r.queries.GetUserByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by id: %w", err)
	}

	credential, err := r.queries.GetCredentialByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get credential by id: %w", err)
	}

	profile, err := r.queries.GetProfileByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get profile by id: %w", err)
	}

	return &identity.IdentityAggregate{
		User:       *toDomainUser(user),
		Credential: toDomainCredential(credential),
		Profile:    toDomainProfile(profile),
	}, nil
}

func (r *userRepository) FindByUserID(ctx context.Context, id string) (*identity.User, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, nil
	}

	pgUUid := common.ToPgUUID(uuid)

	user, err := r.queries.GetUserByID(ctx, pgUUid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra get user by id: %w", err)
	}

	return toDomainUser(user), nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*identity.User, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by email: %w", err)
	}

	return toDomainUser(user), nil
}

func (r *userRepository) FindForLogin(ctx context.Context, email string) (*identity.IdentityAggregate, error) {
	user, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user by id: %w", err)
	}

	credential, err := r.queries.GetCredentialByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get credential by id: %w", err)
	}

	return &identity.IdentityAggregate{
		User:       *toDomainUser(user),
		Credential: toDomainCredential(credential),
		Profile:    nil,
	}, nil
}

func (r *userRepository) FindLoginStatByID(ctx context.Context, id string) (*identity.LoginStat, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, nil
	}

	pgUUid := common.ToPgUUID(uuid)

	stats, err := r.queries.GetLoginStatsByID(ctx, pgUUid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get login stats by id: %w", err)
	}

	return toDomainLoginStat(stats), nil
}
