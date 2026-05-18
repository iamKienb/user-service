package verify_permission

import "context"

type service interface {
	VerifyPermission(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service service
}

func NewHandler(service service) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.VerifyPermission(ctx, cmd)
}
