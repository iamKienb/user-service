package account

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type GenderEnum string

const (
	GenderMale   GenderEnum = "MALE"
	GenderFemale GenderEnum = "FEMALE"
	GenderOther  GenderEnum = "OTHER"
)

type UserProfile struct {
	UserID      shared.UserID
	FullName    string
	Gender      GenderEnum
	PhoneNumber *string
	AvatarURL   *string
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

func (e GenderEnum) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOther:
		return true
	}

	return false
}
