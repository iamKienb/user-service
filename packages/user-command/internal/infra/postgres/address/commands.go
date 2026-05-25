package address

import (
	"context"
	"fmt"
	domain_address "user-command-module/internal/domain/address"
)

func (r *userAddressRepository) CreateUserAddress(ctx context.Context, addr *domain_address.UserAddress) error {
	if err := r.getQuerier(ctx).SaveUserAddress(ctx, toInfraUserAddress(addr)); err != nil {
		return fmt.Errorf("infra: save user address failed: %w", err)
	}

	return nil
}
