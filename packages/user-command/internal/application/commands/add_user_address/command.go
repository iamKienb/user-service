package add_user_address

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type LocationInfo struct {
	ID   string
	Name string
}

type Command struct {
	UserID shared.UserID

	Country  LocationInfo
	Province LocationInfo
	Ward     LocationInfo

	AddressLine  string
	ReceiverName string
	PhoneNumber  string

	Label     string
	IsDefault bool
}

type Result struct {
	UserAddressID shared.UserAddressID
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
