package address

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type UserAddressAddedEvent struct {
	UserID        shared.UserID
	UserAddressID shared.UserAddressID

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
	CreatedAt time.Time
}

func (e UserAddressAddedEvent) EventName() string {
	return "user-service.user.address.added"
}

func (e UserAddressAddedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"user_id":         e.UserID.String(),
		"user_address_id": e.UserAddressID.String(),
		"country_id":      e.CountryID,
		"country_name":    e.CountryName,
		"city_id":         e.CityID,
		"city_name":       e.CityName,
		"district_id":     e.DistrictID,
		"district_name":   e.DistrictName,
		"ward_id":         e.WardID,
		"ward_name":       e.WardName,
		"address_line":    e.AddressLine,
		"receiver_name":   e.ReceiverName,
		"phone_number":    e.PhoneNumber,
		"label":           string(e.Label),
		"isDefault":       e.IsDefault,
		"created_at":      e.CreatedAt,
	}
}
