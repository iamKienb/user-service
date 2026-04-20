package identity

import "errors"

var (
	// Email
	ErrEmailEmpty           = errors.New("email_empty")
	ErrEmailInvalid         = errors.New("email_invalid")
	ErrEmailTaken           = errors.New("email_already_taken")
	ErrAccountLocked        = errors.New("account_locked")
	ErrInvalidCredentials   = errors.New("invalid_credentials")
	ErrInvalidEmailPassword = errors.New("invalid_email_password")

	// Name
	ErrNameEmpty   = errors.New("name_empty")
	ErrNameTooLong = errors.New("name_too_long")

	// User & Profile
	ErrUserInvalid        = errors.New("user_invalid")
	ErrUserNotFound       = errors.New("user_not_found")
	ErrUserNotActive      = errors.New("user_not_active")
	ErrProfileNotFound    = errors.New("profile_not_found")
	ErrCredentialNotFound = errors.New("credential_not_found")

	// OTP & Session
	ErrOTPInvalid     = errors.New("otp_invalid")
	ErrOTPExpired     = errors.New("otp_expired")
	ErrSessionInvalid = errors.New("session_invalid")
	ErrOTPMaxRetry    = errors.New("otp_max_retry_reached")
	ErrResendLimit    = errors.New("otp_resend_limit_reached")
)
