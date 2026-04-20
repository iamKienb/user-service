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

func (r *userRepository) Save(ctx context.Context, agg *identity.IdentityAggregate) error {
	q := r.getQuerier(ctx)

	if err := q.InsertUser(ctx, toInfraUser(&agg.User)); err != nil {
		if common.IsDuplicateEmail(err) {
			return identity.ErrEmailTaken
		}
		return fmt.Errorf("infra: insert user failed: %w", err)
	}

	if err := q.InsertCredential(ctx, toInfraCredential(agg.Credential)); err != nil {
		return fmt.Errorf("infra: insert credential failed: %w", err)
	}

	if err := q.InsertProfile(ctx, toInfraProfile(agg.Profile)); err != nil {
		return fmt.Errorf("infra: insert profile failed: %w", err)
	}

	return nil
}

func (r *userRepository) FindAggregateByID(ctx context.Context, id string) (*identity.IdentityAggregate, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, identity.ErrUserInvalid
	}

	pgUUID := common.ToPgUUID(uuid)

	user, err := r.queries.GetUserByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, identity.ErrUserNotFound
		}
		return nil, fmt.Errorf("infra: get user by id: %w", err)
	}

	credential, err := r.queries.GetCredentialByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, identity.ErrCredentialNotFound
		}
		return nil, fmt.Errorf("infra: get credential by id: %w", err)
	}

	profile, err := r.queries.GetProfileByID(ctx, pgUUID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, identity.ErrProfileNotFound
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

func (r *userRepository) UpdateUser(ctx context.Context, user *identity.User) error {
	if err := r.getQuerier(ctx).UpdateUser(ctx, toUpdateUserInfra(user)); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return fmt.Errorf("infra: update user: %w", err)
	}

	return nil
}

// func (r *postgresUserRepo) FindForLogin(
// 	ctx context.Context,
// 	email string,
// ) (*user.User, *user.Credential, error) {
// 	row, err := r.getQuerier(ctx).GetUserForLogin(ctx, email)
// 	if err != nil {
// 		if errors.Is(err, pgx.ErrNoRows) {
// 			// trả ErrInvalidCredentials thay vì ErrNotFound
// 			// attacker không biết email có tồn tại hay không
// 			return nil, nil, user.ErrInvalidCredentials
// 		}
// 		return nil, nil, fmt.Errorf("find for login: %w", err)
// 	}

// 	return toDomainUser(row.User), toDomainCredential(row.UserCredential), nil
// }
