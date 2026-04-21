package login_user

import (
	"context"
	"time"
)

type Command struct {
	Email    string
	Password string
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
