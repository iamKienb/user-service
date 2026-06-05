package models

import "encoding/json"

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
	Addresses []UserAddress `json:"address"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type userAlias User
	var raw struct {
		userAlias
		Address json.RawMessage `json:"address"`
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	*u = User(raw.userAlias)
	u.Addresses = decodeOneOrMany[UserAddress](raw.Address)
	return nil
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
	FullAddress  string      `json:"full_address"`
	Country      LocationRef `json:"country"`
	Province     LocationRef `json:"province"`
	District     LocationRef `json:"district"`
	Ward         LocationRef `json:"ward"`
	AddressLine  string      `json:"address_line"`
	ReceiverName string      `json:"receiver_name"`
	PhoneNumber  string      `json:"phone_number"`
	Label        string      `json:"label"`
	IsDefault    bool        `json:"is_default"`
}

type UserPage struct {
	Items         []User
	Total         int64
	NextPageToken string
}

func decodeOneOrMany[T any](raw json.RawMessage) []T {
	if len(raw) == 0 || string(raw) == "null" {
		return nil
	}
	var many []T
	if err := json.Unmarshal(raw, &many); err == nil {
		return many
	}
	var one T
	if err := json.Unmarshal(raw, &one); err != nil {
		return nil
	}
	return []T{one}
}
