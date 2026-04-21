package shared

import "time"

const (
	EmailCacheTTL = 30 * time.Minute

	OTPMaxRetry = 5
	OTPTTL      = 2 * time.Minute

	SessionTTL      = 10 * time.Minute
	ResendMaxCount  = 5
	ResendWindowTTL = 5 * time.Minute
)
