package profile

import (
	"context"
	"fmt"
	domain_profile "user-command-module/internal/domain/profile"
)

func (r *profileRepository) CreateProfile(ctx context.Context, profile *domain_profile.Profile) error {
	if err := r.getQuerier(ctx).CreateUserProfile(ctx, toInfraProfile(profile)); err != nil {
		return fmt.Errorf("infra: save profile failed: %w", err)
	}

	return nil
}
