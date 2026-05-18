package assign_member

import "context"

type shopService interface {
	AssignMember(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service shopService
}

func NewHandler(service shopService) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.AssignMember(ctx, cmd)
}
