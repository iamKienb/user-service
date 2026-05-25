package events

import "time"

const TopicUserAddressAdded = "user-service.user.address.added"

type UserAddressAdded struct {
	UserAddressID string `json:"user_address_id"`
	UserID        string `json:"user_id"`

	CountryID   int    `json:"country_id"`
	CountryName string `json:"country_name"`

	CityID   int    `json:"city_id"`
	CityName string `json:"city_name"`

	DistrictID   int    `json:"district_id"`
	DistrictName string `json:"district_name"`

	WardID   int    `json:"ward_id"`
	WardName string `json:"ward_name"`

	AddressLine  string `json:"address_line"`
	ReceiverName string `json:"receiver_name"`
	PhoneNumber  string `json:"phone_number"`
	Label        string `json:"label"`
	IsDefault    bool   `json:"is_default"`

	CreatedAt time.Time `json:"created_at"`
}
