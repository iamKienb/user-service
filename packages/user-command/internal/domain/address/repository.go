package address

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	FindAddressByID(ctx context.Context, userAddressID shared.UserAddressID) (*UserAddress, error)
}

type CommandRepository interface {
	CreateUserAddress(ctx context.Context, address *UserAddress) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
