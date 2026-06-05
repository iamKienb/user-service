package get_user_detail

import (
	"context"

	"user-query-module/internal/application/port"
)

type Query struct {
	UserID string
}

type Result struct {
	User *port.User
}

type Executor interface {
	Execute(ctx context.Context, query Query) (*Result, error)
}
