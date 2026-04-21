package resend_otp

import "context"

type otpService interface {
	Resend(ctx context.Context, cmd Command) (*Result, error)
}

type handler struct {
	service otpService
}

func NewHandler(service otpService) Executor {
	return &handler{
		service: service,
	}
}

func (h *handler) Execute(ctx context.Context, cmd Command) (*Result, error) {
	return h.service.Resend(ctx, cmd)
}
