package events

import "time"

const TopicUserAddressAdded = "user-service.user.address.added"

type UserAddressAdded struct {
	UserAddressID string `json:"user_address_id"`
	UserID        string `json:"user_id"`

	CountryID   string `json:"country_id"`
	CountryName string `json:"country_name"`

	ProvinceID   string `json:"province_id"`
	ProvinceName string `json:"province_name"`

	WardID   string `json:"ward_id"`
	WardName string `json:"ward_name"`

	AddressLine  string `json:"address_line"`
	ReceiverName string `json:"receiver_name"`
	PhoneNumber  string `json:"phone_number"`
	Label        string `json:"label"`
	IsDefault    bool   `json:"is_default"`

	CreatedAt time.Time `json:"created_at"`
}
