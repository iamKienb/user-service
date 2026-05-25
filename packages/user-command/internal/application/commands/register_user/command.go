package register_user

import (
	"context"
	"time"
)

type UserProfile struct {
	FullName string
	Gender   string
}
type Command struct {
	Email    string
	Password string
	Profile  UserProfile
}

type Result struct {
	SessionToken string
	ExpiresAt    time.Time
}

type Executor interface {
	Execute(ctx context.Context, cmd Command) (*Result, error)
}
