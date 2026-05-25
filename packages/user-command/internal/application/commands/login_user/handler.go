package login_user

import "context"

type userService interface {
	Login(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service userService
}

func NewHandler(service userService) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.Login(ctx, cmd)
}
