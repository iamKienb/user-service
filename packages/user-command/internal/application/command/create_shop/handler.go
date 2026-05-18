package create_shop

import "context"

type shopService interface {
	CreateShop(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service shopService
}

func NewHandler(service shopService) Executor {
	return &handler{service: service}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.CreateShop(ctx, cmd)
}
