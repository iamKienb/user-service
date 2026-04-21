package user

import (
	"context"
	"fmt"
	"shopify-user-command-module/internal/domain/identity"
	"shopify-user-command-module/internal/infra/common"
)

func (r *userRepository) SaveAggregate(ctx context.Context, agg *identity.IdentityAggregate) error {
	q := r.getQuerier(ctx)

	if err := q.InsertUser(ctx, toInfraUser(&agg.User)); err != nil {
		if common.IsDuplicateEmail(err) {
			return nil
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

func (r *userRepository) UpdateUser(ctx context.Context, user *identity.User) error {
	if err := r.getQuerier(ctx).UpdateUser(ctx, toUpdateUserInfra(user)); err != nil {
		return fmt.Errorf("infra: update user: %w", err)
	}

	return nil
}

func (r *userRepository) SaveLoginStat(ctx context.Context, loginStat *identity.LoginStat) error {
	if err := r.getQuerier(ctx).SaveLoginStats(ctx, toInfraLoginStat(loginStat)); err != nil {
		return fmt.Errorf("infra: save login stats: %w", err)
	}

	return nil
}
