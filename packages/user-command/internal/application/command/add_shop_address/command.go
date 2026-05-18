package add_shop_address

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type LocationInfo struct {
	ID   int
	Name string
}

type Command struct {
	ShopID shared.ShopID
	UserID shared.UserID

	Country  LocationInfo
	City     LocationInfo
	District LocationInfo
	Ward     LocationInfo

	AddressLine string
	ContactName string
	PhoneNumber string

	Type string
}

type Result struct {
	ShopAddressID string
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
