package add_user_address

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type LocationInfo struct {
	ID   int
	Name string
}

type Command struct {
	UserID shared.UserID

	Country  LocationInfo
	City     LocationInfo
	District LocationInfo
	Ward     LocationInfo

	AddressLine  string
	ReceiverName string
	PhoneNumber  string

	Label     string
	IsDefault bool
}

type Result struct {
	UserAddressID string
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
