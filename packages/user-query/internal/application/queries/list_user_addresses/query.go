package list_user_addresses

import (
	"context"

	"user-query-module/internal/application/port"
)

type Query struct {
	UserID string
}

type Result struct {
	Addresses []port.UserAddress
}

type Executor interface {
	Execute(ctx context.Context, query Query) (*Result, error)
}
