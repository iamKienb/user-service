package search_users

import "context"

type userQueryService interface {
	SearchUsers(ctx context.Context, query Query) (*Result, error)
}

type handler struct {
	service userQueryService
}

func NewHandler(service userQueryService) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, query Query) (*Result, error) {
	return h.service.SearchUsers(ctx, query)
}
