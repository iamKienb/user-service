package get_user_address_by_id

import (
	"context"
)

type userService interface {
	GetAddress(ctx context.Context, qry Query) (*Result, error)
}

type handler struct {
	service userService
}

func NewHandler(service userService) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, qry Query) (*Result, error) {
	return h.service.GetAddress(ctx, qry)
}
