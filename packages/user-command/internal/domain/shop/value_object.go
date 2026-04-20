package user

import (
	"errors"
	"net/mail"
	"strings"
	"unicode"

	"github.com/google/uuid"
)

type UserID uuid.UUID

func NewUserID() UserID {
	return UserID(uuid.Must(uuid.NewV7()))
}

type Email string

func NewEmail(email string) (Email, error) {
	if email == "" {
		return Email(""), errors.New("email cannot be empty")
	}
	address, err := mail.ParseAddress(email)
	if err != nil {
		return Email(""), errors.New("invalid email format")
	}

	return Email(address.Address), nil
}

type Password string

func NewPassword(password string) (Password, error) {
	if len(password) < 8 {
		return Password(""), errors.New("password must be at least 8 characters long")
	}

	var hasUpper, hasNumber bool

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsNumber(char):
			hasNumber = true
		}
	}

	if !hasUpper || !hasNumber {
		return Password(""), errors.New("password must include at least one uppercase letter and one number")
	}

	return Password(password), nil
}

type Name string

func NewName(raw string) (Name, error) {
	name := strings.TrimSpace(raw)

	if len(name) == 0 {
		return Name(""), errors.New("full name cannot be empty")
	}
	if len(name) > 100 {
		return Name(""), errors.New("full name too long, max 100 characters")
	}
	return Name(name), nil
}
