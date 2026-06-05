package address

import (
	"user-command-module/internal/domain/shared"
)

type NewUserAddressParams struct {
	UserID shared.UserID

	CountryID   string
	CountryName string

	ProvinceID   string
	ProvinceName string

	WardID   string
	WardName string

	AddressLine  string
	ReceiverName string
	PhoneNumber  string

	Label     LabelEnum
	IsDefault bool
}
