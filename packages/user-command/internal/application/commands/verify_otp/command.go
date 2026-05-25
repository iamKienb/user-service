package verify_otp

import (
	"context"
	"time"
)

type Command struct {
	OTP          string
	SessionToken string
}

type Result struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
