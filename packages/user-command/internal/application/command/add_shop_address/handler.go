package add_shop_address

import "context"

type UserService interface {
	AddAddress(ctx context.Context, cmd Command) (*Result, error)
}

type Handler struct {
	service UserService
}

func NewHandler(service UserService) Executor {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.AddAddress(ctx, cmd)
}
