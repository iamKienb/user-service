package address

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	FindUserAddressByID(ctx context.Context, addressID shared.UserAddressID) (*UserAddress, error)
}

type CommandRepository interface {
	CreateUserAddress(ctx context.Context, address *UserAddress) error
}

type Repository interface {
	// QueryRepository
	CommandRepository
}
