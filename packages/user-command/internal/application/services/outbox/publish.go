package outbox

import (
	"context"
	"time"
	"user-command-module/internal/application/port"

	"github.com/google/uuid"
)

func (s *outboxService) PublishBatch(ctx context.Context, params []port.OutboxParam) error {
	if len(params) == 0 {
		return nil
	}

	totalEvents := 0
	for _, param := range params {
		totalEvents += len(param.Events)
	}
	if totalEvents == 0 {
		return nil
	}

	messages := make([]port.OutboxEvent, 0, totalEvents)
	now := time.Now().UTC()

	for _, param := range params {
		for _, event := range param.Events {
			messages = append(messages, port.OutboxEvent{
				ID:             uuid.New(),
				AggregateID:    param.AggregateID,
				AggregateType:  param.AggregateType,
				EventType:      event.EventName(),
				Payload:        event.IntegrationPayload(),
				PartitionKey:   param.AggregateID.String(),
				IdempotencyKey: uuid.New(),
				CreatedAt:      now,
			})
		}
	}

	return s.outboxRepo.SaveOutboxBatch(ctx, messages)
}
