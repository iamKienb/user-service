package verify_otp

import "context"

type otpService interface {
	Verify(ctx context.Context, cmd Command) (*Result, error)
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
	return h.service.Verify(ctx, cmd)
}
