package user

import "errors"

var (
	ErrEmailEmpty   = errors.New("email_empty")
	ErrEmailInvalid = errors.New("email_invalid")
	ErrEmailTaken   = errors.New("email_already_taken")

	ErrUserInvalid   = errors.New("user_invalid")
	ErrUserNotFound  = errors.New("user_not_found")
	ErrUserNotActive = errors.New("user_not_active")

	ErrCredentialNotFound = errors.New("credential_not_found")
)
