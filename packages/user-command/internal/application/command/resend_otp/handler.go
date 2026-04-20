package resend_otp

import "context"

type otpService interface {
	Resend(ctx context.Context, cmd Command) (*Result, error)
}

type Handler struct {
	service otpService
}

func NewHandler(service otpService) Executor {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	result, err := h.service.Resend(ctx, cmd)
	if err != nil {
		return nil, err
	}
	return result, nil
}
