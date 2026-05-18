package shared

import "time"

const (
	EmailCacheTTL = 30 * time.Minute

	OTPMaxRetry = 5
	OTPTTL      = 10 * time.Minute

	SessionTTL      = 60 * time.Minute
	ResendMaxCount  = 5
	ResendWindowTTL = 5 * time.Minute

	IdemKeyTTL = 5 * time.Second
)
