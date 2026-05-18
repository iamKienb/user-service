package create_shop

import "context"

type service interface {
	CreateShop(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service service
}

func NewHandler(service service) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.CreateShop(ctx, cmd)
}
