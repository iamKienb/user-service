package account

import (
	"strings"
	"time"
)

type Gender string

const (
	MALE   Gender = "MALE"
	FEMALE Gender = "FEMALE"
	OTHER  Gender = "OTHER"
)

type Profile struct {
	UserID      UserID
	FullName    string
	Gender      Gender
	PhoneNumber *string
	AvatarURL   *string
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (g Gender) IsValid() bool {
	switch g {
	case MALE, FEMALE, OTHER:
		return true
	}
	return false
}

func ValidateGender(gender string) (Gender, error) {
	normalized := Gender(strings.ToUpper(strings.TrimSpace(gender)))
	if !normalized.IsValid() {
		return "", ErrGenderInvalid
	}

	return normalized, nil
}
