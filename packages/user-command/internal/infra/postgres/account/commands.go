package account

import (
	"context"
	"fmt"

	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/infra/common"
)

func (r *accountRepository) SaveAggregate(ctx context.Context, agg *account.Aggregate) error {
	q := r.getQuerier(ctx)

	if err := q.InsertUser(ctx, toInfraUser(&agg.User)); err != nil {
		if common.IsDuplicateEmail(err) {
			return account.ErrEmailTaken
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

func (r *accountRepository) UpdateUser(ctx context.Context, user *account.User) error {
	if err := r.getQuerier(ctx).UpdateUser(ctx, toUpdateUserInfra(user)); err != nil {
		return fmt.Errorf("infra: update user: %w", err)
	}

	return nil
}
