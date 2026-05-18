package account

import "errors"

var (
	ErrEmailEmpty      = errors.New("email_empty")
	ErrEmailInvalid    = errors.New("email_invalid")
	ErrEmailTaken      = errors.New("email_already_taken")
	ErrNameEmpty       = errors.New("name_empty")
	ErrNameTooLong     = errors.New("name_too_long")
	ErrGenderInvalid   = errors.New("gender_invalid")
	ErrUserInvalid     = errors.New("user_invalid")
	ErrUserNotFound    = errors.New("user_not_found")
	ErrUserNotActive   = errors.New("user_not_active")
	ErrProfileNotFound = errors.New("profile_not_found")

	ErrCredentialNotFound = errors.New("credential_not_found")

	ErrMaxAddressReached = errors.New("only had create 5 address")
)
