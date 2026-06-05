package get_user_detail

import "context"

type userQueryService interface {
	GetUserDetail(ctx context.Context, query Query) (*Result, error)
}

type handler struct {
	service userQueryService
}

func NewHandler(service userQueryService) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, query Query) (*Result, error) {
	return h.service.GetUserDetail(ctx, query)
}
