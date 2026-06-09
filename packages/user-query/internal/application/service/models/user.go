package models

type Page struct {
	Size  int
	Token string
}

type User struct {
	ID        string        `json:"id"`
	Email     string        `json:"email"`
	Status    string        `json:"status"`
	Roles     []string      `json:"roles"`
	Profile   *UserProfile  `json:"profile"`
	Addresses []UserAddress `json:"addresses"`
}

type UserProfile struct {
	FullName string `json:"full_name"`
	Gender   string `json:"gender"`
}

type LocationRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserAddress struct {
	ID           string      `json:"id"`
	UserID       string      `json:"user_id"`
	Country      LocationRef `json:"country"`
	Province     LocationRef `json:"province"`
	Ward         LocationRef `json:"ward"`
	AddressLine  string      `json:"address_line"`
	FullAddress  string      `json:"full_address"`
	ReceiverName string      `json:"receiver_name"`
	PhoneNumber  string      `json:"phone_number"`
	Label        string      `json:"label"`
	IsDefault    bool        `json:"is_default"`
	CreatedAt    string      `json:"created_at"`
	UpdatedAt    string      `json:"updated_at"`
}

type UserPage struct {
	Items         []User
	Total         int64
	NextPageToken string
}
