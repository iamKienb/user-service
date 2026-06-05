package port

import "encoding/json"

type Page struct {
	Size  int
	Token string
}

type User struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Status    string         `json:"status"`
	Roles     []string       `json:"roles"`
	Profile   *UserProfile   `json:"profile"`
	Addresses []UserAddress  `json:"address"`
	Extra     map[string]any `json:"-"`
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
	if len(raw.Address) == 0 || string(raw.Address) == "null" {
		return nil
	}

	var addresses []UserAddress
	if err := json.Unmarshal(raw.Address, &addresses); err == nil {
		u.Addresses = addresses
		return nil
	}

	var address UserAddress
	if err := json.Unmarshal(raw.Address, &address); err != nil {
		return err
	}
	u.Addresses = []UserAddress{address}
	return nil
}

type UserProfile struct {
	FullName string         `json:"full_name"`
	Gender   string         `json:"gender"`
	Extra    map[string]any `json:"-"`
}

type UserAddress struct {
	ID           string         `json:"id"`
	UserID       string         `json:"user_id"`
	FullAddress  string         `json:"full_address"`
	AddressLine  string         `json:"address_line"`
	ReceiverName string         `json:"receiver_name"`
	PhoneNumber  string         `json:"phone_number"`
	Label        string         `json:"label"`
	IsDefault    bool           `json:"is_default"`
	Extra        map[string]any `json:"-"`
}

type UserPage struct {
	Items         []User
	Total         int64
	NextPageToken string
}
