package address

import "time"

type Address struct {
	ID     string
	UserID string

	Label        string
	ReceiverName string
	PhoneNumber  string

	AddressLine string
	Ward        string
	District    string
	City        string
	Country     string

	IsDefault bool

	CreatedAt time.Time
	UpdatedAt time.Time
}
