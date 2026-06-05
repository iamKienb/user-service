package user

import (
	"context"
	"fmt"
	domain_user "user-command-module/internal/domain/user"
)

func (r *userRepository) CreateUser(ctx context.Context, user *domain_user.User) error {
	q := r.getQuerier(ctx)

	if err := q.CreateUser(ctx, toInfraUser(user)); err != nil {
		if r.IsDuplicateEmail(err) {
			return domain_user.ErrEmailTaken
		}
		return fmt.Errorf("infra: save user failed: %w", err)
	}

	if err := q.CreateUserCredential(ctx, toInfraCredential(&user.Credential)); err != nil {
		return fmt.Errorf("infra: save credential failed: %w", err)
	}

	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *domain_user.User) error {
	if err := r.getQuerier(ctx).UpdateUser(ctx, toUpdateUserInfra(user)); err != nil {
		return fmt.Errorf("infra: update user: %w", err)
	}

	return nil
}
