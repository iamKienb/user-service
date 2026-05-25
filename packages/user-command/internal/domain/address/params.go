package address

import (
	"user-command-module/internal/domain/shared"
)

type NewUserAddressParams struct {
	UserID shared.UserID

	CountryID   int
	CountryName string

	CityID   int
	CityName string

	DistrictID   int
	DistrictName string

	WardID   int
	WardName string

	AddressLine  string
	ReceiverName string
	PhoneNumber  string

	Label     LabelEnum
	IsDefault bool
}
