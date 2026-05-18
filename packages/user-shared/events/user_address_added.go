package events

import "time"

const TopicUserAddressAdded = "user-service.user.address.added"

type UserAddressAdded struct {
	UserAddressID string `json:"user_address_id"`
	UserID        string `json:"user_id"`

	CountryID   int    `json:"country_id"`
	CountryName string `json:"country_name"`

	CityID   int
	CityName string

	DistrictID   int
	DistrictName string

	WardID   int
	WardName string

	AddressLine  string
	ReceiverName string
	PhoneNumber  string
	Label        string
	IsDefault    bool

	created_at time.Time
	updated_at time.Time
}

func (u UserAddressAdded) EventName() string {
	return TopicUserAddressAdded
}
