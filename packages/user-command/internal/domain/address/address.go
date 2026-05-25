package address

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type LabelEnum string

const (
	LabelOffice LabelEnum = "OFFICE"
	LabelHouse  LabelEnum = "HOUSE"
)

type UserAddress struct {
	ID     shared.UserAddressID
	UserID shared.UserID

	CountryID  int
	CityID     int
	DistrictID int
	WardID     int

	AddressLine  string
	ReceiverName string
	PhoneNumber  string

	Label     LabelEnum
	IsDefault bool

	CreatedAt time.Time
	UpdatedAt *time.Time
	shared.EventEntity
}

func NewUserAddress(params NewUserAddressParams) *UserAddress {
	now := time.Now().UTC()
	userAddressID := shared.NewID[shared.UserAddressID]()

	address := &UserAddress{
		ID:     userAddressID,
		UserID: params.UserID,

		CountryID:  params.CountryID,
		CityID:     params.CityID,
		DistrictID: params.DistrictID,
		WardID:     params.WardID,

		AddressLine:  params.AddressLine,
		ReceiverName: params.ReceiverName,
		PhoneNumber:  params.PhoneNumber,
		Label:        params.Label,
		IsDefault:    params.IsDefault,

		CreatedAt: now,
		UpdatedAt: nil,
	}

	address.AddEvent(UserAddressAddedEvent{
		UserAddressID: userAddressID,
		UserID:        address.UserID,

		CountryID:   params.CityID,
		CountryName: params.CountryName,

		CityID:   params.CityID,
		CityName: params.CityName,

		DistrictID:   params.DistrictID,
		DistrictName: params.DistrictName,

		WardID:   params.WardID,
		WardName: params.WardName,

		AddressLine:  address.AddressLine,
		ReceiverName: address.ReceiverName,
		PhoneNumber:  address.PhoneNumber,
		Label:        address.Label,
		IsDefault:    address.IsDefault,
		CreatedAt:    address.CreatedAt,
	})

	return address
}

func (e LabelEnum) IsValid() bool {
	switch e {
	case LabelOffice, LabelHouse:
		return true
	}

	return false
}
