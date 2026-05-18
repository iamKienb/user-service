package add_user_address

import "context"

type service interface {
	AddAddress(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service service
}

func NewHandler(service service) Executor {
	return &handler{
		service: service,
	}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.AddAddress(ctx, cmd)
}
