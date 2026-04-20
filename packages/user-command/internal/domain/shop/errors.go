package user

import "errors"

var (
	ErrEmailTaken = errors.New("email already taken")

	ErrNotFound = errors.New("user not found")

	ErrAccountLocked = errors.New("account is temporarily locked")

	ErrInvalidCredentials = errors.New("invalid credentials")
)
