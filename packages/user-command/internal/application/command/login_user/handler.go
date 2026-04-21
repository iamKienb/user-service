package login_user

import (
	"context"
	"fmt"
)

type userService interface {
	Login(ctx context.Context, cmd Command) (*Result, error)
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
	fmt.Println("CMD1", cmd)

	result, err := h.service.Login(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return result, nil
}
