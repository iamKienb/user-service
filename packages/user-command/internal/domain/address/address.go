package address

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type LabelEnum string

const (
	LabelOffice LabelEnum = "OFFICE"
	LabelHouse  LabelEnum = "HOME"
)

type UserAddress struct {
	ID     shared.UserAddressID
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

		CountryID:   params.CountryID,
		CountryName: params.CountryName,

		ProvinceID:   params.ProvinceID,
		ProvinceName: params.ProvinceName,
		WardID:       params.WardID,
		WardName:     params.WardName,

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

		CountryID:   params.CountryID,
		CountryName: params.CountryName,

		ProvinceID:   params.ProvinceID,
		ProvinceName: params.ProvinceName,

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

func (a *UserAddress) FlushEvents() []shared.DomainEvent {
	events := a.Flush()

	return events
}

func (a *UserAddress) Type() string {
	return "USER_ADDRESS"
}

func (e LabelEnum) IsValid() bool {
	switch e {
	case LabelOffice, LabelHouse:
		return true
	}

	return false
}
