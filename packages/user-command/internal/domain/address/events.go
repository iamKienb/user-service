package address

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type UserAddressAddedEvent struct {
	UserID        shared.UserID
	UserAddressID shared.UserAddressID

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
		"province_id":     e.ProvinceID,
		"province_name":   e.ProvinceName,
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
