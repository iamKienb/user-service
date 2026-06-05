package get_user_profile

import (
	"context"

	"user-query-module/internal/application/port"
)

type Query struct {
	UserID string
}

type Result struct {
	Profile *port.UserProfile
}

type Executor interface {
	Execute(ctx context.Context, query Query) (*Result, error)
}
