package auth

import "errors"

var (
	ErrAccountLocked      = errors.New("account_locked")
	ErrInvalidCredentials = errors.New("invalid_credentials")

	ErrOTPInvalid     = errors.New("otp_invalid")
	ErrOTPExpired     = errors.New("otp_expired")
	ErrSessionInvalid = errors.New("session_invalid")
	ErrOTPMaxRetry    = errors.New("otp_max_retry_reached")
	ErrResendLimit    = errors.New("otp_resend_limit_reached")

	ErrActionNotDefined = errors.New("action_not_defined")
	ErrShopDenied       = errors.New("shop_permission_denied")
	ErrProductDenied    = errors.New("product_permission_denied")
	ErrInventoryDenied  = errors.New("inventory_permission_denied")
)
