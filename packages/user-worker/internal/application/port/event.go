package port

import (
	"context"
	"encoding/json"
	kafkax "user-worker-module/internal/infra/kafka"
)

type Message = kafkax.Message

type EventProcessor interface {
	Handle(ctx context.Context, msg Message) error
}

type EventHandler interface {
	Handle(ctx context.Context, payload json.RawMessage) error
}
