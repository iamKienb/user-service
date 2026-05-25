package add_user_address

import "context"

type userService interface {
	AddAddress(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service userService
}

func NewHandler(service userService) Executor {
	return &handler{
		service: service,
	}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.AddAddress(ctx, cmd)
}
