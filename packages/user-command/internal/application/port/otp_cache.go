package port

import (
	"context"
	"time"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
)

type OTPEntry struct {
	OTP       string
	UserID    string
	Email     string
	ExpiresAt time.Time
	FailCount int
}

type SessionEntry struct {
	UserID      string
	Email       string
	ResendCount int
}

type OTPCache interface {
	redisx.Cache
	SaveOTP(ctx context.Context, sessionToken string, entry OTPEntry, ttl time.Duration) error
	GetOTP(ctx context.Context, sessionToken string) (*OTPEntry, error)
	DeleteOTP(ctx context.Context, sessionToken string) error

	SaveSession(ctx context.Context, sessionToken string, session SessionEntry, ttl time.Duration) error
	GetSession(ctx context.Context, sessionToken string) (*SessionEntry, error)
	DeleteSession(ctx context.Context, sessionToken string) error
	IncrResendCount(ctx context.Context, sessionToken string, ttl time.Duration) (int64, error)
}
