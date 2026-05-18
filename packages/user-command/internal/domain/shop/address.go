package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type AddressTypeEnum string

const (
	TypePickup AddressTypeEnum = "PICKUP"
	TypeReturn AddressTypeEnum = "RETURN"
)

type ShopAddress struct {
	ID     shared.ShopAddressID
	ShopID shared.ShopID

	CountryID  int
	CityID     int
	DistrictID int
	WardID     int

	ContactName string
	PhoneNumber string
	AddressLine string
	Type        AddressTypeEnum

	CreatedBy *shared.UserID
	UpdatedBy *shared.UserID

	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewShopAddress(p AddressParams) *ShopAddress {
	now := time.Now().UTC()
	addressID := shared.NewID[shared.ShopAddressID]()

	return &ShopAddress{
		ID:     addressID,
		ShopID: p.ShopID,

		CountryID:  p.CountryID,
		CityID:     p.CityID,
		DistrictID: p.DistrictID,
		WardID:     p.WardID,

		Type:        p.Type,
		ContactName: p.ContactName,
		PhoneNumber: p.PhoneNumber,
		AddressLine: p.AddressLine,

		CreatedBy: &p.UserID,
		UpdatedBy: &p.UserID,

		CreatedAt: now,
		UpdatedAt: nil,
	}
}

func (e AddressTypeEnum) IsValidType() bool {
	switch e {
	case TypePickup, TypeReturn:
		return true
	}

	return false
}

func (e AddressTypeEnum) IsValid() bool {
	return e.IsValidType()
}
