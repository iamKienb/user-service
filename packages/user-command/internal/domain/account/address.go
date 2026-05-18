package account

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type LabelEnum string

const (
	LabelPickup LabelEnum = "PICKUP"
	LabelReturn LabelEnum = "RETURN"
)

func (e LabelEnum) IsValid() bool {
	switch e {
	case LabelPickup, LabelReturn:
		return true
	}

	return false
}

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
	UpdatedAt time.Time
	events    shared.EventEntity
}

func NewUserAddress(p AddAddressParams) *UserAddress {
	now := time.Now().UTC()
	userAddressID := shared.NewID[shared.UserAddressID]()

	addr := &UserAddress{
		ID:     userAddressID,
		UserID: p.UserID,

		CountryID:  p.CountryID,
		CityID:     p.CityID,
		DistrictID: p.DistrictID,
		WardID:     p.WardID,

		AddressLine:  p.AddressLine,
		ReceiverName: p.ReceiverName,
		PhoneNumber:  p.PhoneNumber,
		Label:        p.Label,
		IsDefault:    p.IsDefault,

		CreatedAt: now,
		UpdatedAt: now,
	}

	// addr.events.AddEvent(events.UserAddressAdded{
	// 	UserAddressID: userAddressID.String(),
	// 	UserID:        p.UserID.String(),

	// 	CountryID:   p.CityID,
	// 	CountryName: p.CountryName,

	// 	CityID:   p.CityID,
	// 	CityName: p.CityName,

	// 	DistrictID:   p.DistrictID,
	// 	DistrictName: p.DistrictName,

	// 	WardID:   p.WardID,
	// 	WardName: p.WardName,

	// 	AddressLine:  p.AddressLine,
	// 	ReceiverName: p.ReceiverName,
	// 	PhoneNumber:  p.PhoneNumber,
	// 	Label:        string(p.Label),
	// 	IsDefault:    p.IsDefault,
	// })

	return addr
}
