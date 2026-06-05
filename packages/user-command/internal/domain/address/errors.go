package address

import "errors"

var (
	ErrLabelInvalid     = errors.New("invalid_label")
	ErrInvalidAddressID = errors.New("user_address_id_invalid")
	ErrAddressNotFound  = errors.New("address_not_found")
)
