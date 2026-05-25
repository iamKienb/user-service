package profile

import "errors"

var (
	ErrNameEmpty   = errors.New("name_empty")
	ErrNameTooLong = errors.New("name_too_long")

	ErrGenderInvalid = errors.New("gender_invalid")

	ErrProfileNotFound = errors.New("profile_not_found")
)
