package search_users

import (
	"context"

	"user-query-module/internal/application/port"
)

type Query struct {
	Keyword string
	Status  string
	Page    port.Page
}

type Result = port.UserPage

type Executor interface {
	Execute(ctx context.Context, query Query) (*Result, error)
}
