package shop

import "user-command-module/internal/domain/shared"

type AddressParams struct {
	UserID     shared.UserID
	ShopID     shared.ShopID
	CountryID  int
	CityID     int
	DistrictID int
	WardID     int

	Type        AddressTypeEnum
	ContactName string
	PhoneNumber string
	AddressLine string
}

type AggregateParams struct {
	UserID shared.UserID
	Name   string
	Slug   string

	Description *string
	LogoUrl     *string
	BannerUrl   *string
}

type MemberAggregateParams struct {
	ShopID shared.ShopID

	MemberID   shared.UserID
	MemberName string

	AddedBy     shared.UserID
	NameAddedBy string
	RoleIDs     []shared.RoleID
}
