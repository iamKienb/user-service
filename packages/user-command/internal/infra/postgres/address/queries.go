package address

import (
	"context"
	"errors"
	"fmt"
	domain_address "user-command-module/internal/domain/address"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
	"github.com/jackc/pgx/v5"
)

func (r *userAddressRepository) FindAddressByID(ctx context.Context, userAddressID shared.UserAddressID) (*domain_address.UserAddress, error) {
	profileRow, err := r.getQuerier(ctx).FindUserAddressByID(ctx, conv.UUID(userAddressID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user address by id: %w", err)
	}

	return toDomainUserAddress(profileRow), err
}
