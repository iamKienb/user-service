package register_user

import (
	"context"
	"time"
)

type Command struct {
	Email    string
	Password string
	FullName string
	Gender   string
}

type Result struct {
	SessionToken string
	ExpiresAt    time.Time
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
