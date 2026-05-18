package account

import (
	"time"
	"user-command-module/internal/domain/shared"
)

const DefaultPasswordVersion = 1

type UserCredential struct {
	UserID            shared.UserID
	PasswordHash      string
	PasswordVersion   int
	PasswordUpdatedAt time.Time
}
