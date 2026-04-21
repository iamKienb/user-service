package account

import "time"

const DefaultPasswordVersion = 1

type Credential struct {
	UserID            UserID
	PasswordHash      string
	PasswordVersion   int
	PasswordUpdatedAt time.Time
}
