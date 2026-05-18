package account

import (
	"context"
	"fmt"
	"user-command-module/internal/domain/account"
)

func (r *accountRepository) SaveAggregate(ctx context.Context, agg *account.Aggregate) error {
	q := r.getQuerier(ctx)

	if err := q.SaveUser(ctx, toInfraUser(&agg.User)); err != nil {
		if r.IsDuplicateEmail(err) {
			return account.ErrEmailTaken
		}
		return fmt.Errorf("infra: save user failed: %w", err)
	}

	if err := q.SaveUserCredential(ctx, toInfraCredential(agg.Credential)); err != nil {
		return fmt.Errorf("infra: save credential failed: %w", err)
	}

	if err := q.SaveUserProfile(ctx, toInfraProfile(agg.Profile)); err != nil {
		return fmt.Errorf("infra: save profile failed: %w", err)
	}

	return nil
}

func (r *accountRepository) SaveAddress(ctx context.Context, addr *account.UserAddress) error {
	if err := r.getQuerier(ctx).SaveUserAddress(ctx, toInfraUserAddress(addr)); err != nil {
		return fmt.Errorf("infra: save user address failed: %w", err)
	}

	return nil
}

func (r *accountRepository) UpdateUser(ctx context.Context, user *account.User) error {
	if err := r.getQuerier(ctx).UpdateUser(ctx, toUpdateUserInfra(user)); err != nil {
		return fmt.Errorf("infra: update user: %w", err)
	}

	return nil
}
