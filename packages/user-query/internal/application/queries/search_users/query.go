package search_users

import (
	"context"

	"user-query-module/internal/application/service/models"
)

type Query struct {
	Keyword string
	Status  string
	Page    models.Page
}

type Result = models.UserPage

type Executor interface {
	Execute(ctx context.Context, query Query) (*Result, error)
}
