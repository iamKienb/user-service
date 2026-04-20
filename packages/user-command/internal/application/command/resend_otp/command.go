package resend_otp

import (
	"context"
	"time"
)

type Command struct {
	SessionToken string
}

type Result struct {
	ExpiresAt time.Time
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
