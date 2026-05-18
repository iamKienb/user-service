package account

import "user-command-module/internal/domain/shared"

type AggregateParams struct {
	Email        string
	PasswordHash string
	FullName     string
	Gender       GenderEnum
}

type AddAddressParams struct {
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
