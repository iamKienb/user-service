package get_user_address_by_id

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type Query struct {
	UserID        shared.UserID
	UserAddressID shared.UserAddressID
}

type Result struct {
	UserAddressID shared.UserAddressID
	UserID        shared.UserID
	ReceiverName  string
	PhoneNumber   string

	ProvinceID   string
	ProvinceName string

	WardID   string
	WardName string

	AddressLine string
	Label       string
	IsDefault   bool
}

type Executor interface {
	Execute(ctx context.Context, qry Query) (*Result, error)
}
