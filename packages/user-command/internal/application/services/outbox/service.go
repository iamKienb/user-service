package outbox

import (
	"context"
	"user-command-module/internal/application/port"
)

type Service interface {
	Publish(ctx context.Context, param port.OutboxParam) error
}

type outboxService struct {
	outboxRepo port.OutboxRepository
}

func NewOutboxService(outboxRepo port.OutboxRepository) Service {
	return &outboxService{
		outboxRepo: outboxRepo,
	}
}
