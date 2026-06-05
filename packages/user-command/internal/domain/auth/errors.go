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

	ErrAccessDenied = errors.New("access_denied")
)
