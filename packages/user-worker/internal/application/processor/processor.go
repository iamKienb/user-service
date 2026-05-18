package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"user-shared-module/events"
	"user-shared-module/indexing"
	"user-worker-module/internal/application/port"
	"user-worker-module/internal/application/processor/handler"
)

type UserEventProcessor struct {
	handlers map[string]port.EventHandler
}

func NewUserEventProcessor(repo port.ESRepository) port.EventProcessor {
	userAlias := indexing.UserAlias

	p := &UserEventProcessor{
		handlers: make(map[string]port.EventHandler),
	}

	p.handlers[events.TopicUserRegistered] = handler.NewUserRegisterHandler(repo, userAlias)
	p.handlers[events.TopicUserActivated] = handler.NewUserActivatedHandler(repo, userAlias)

	return p
}

func (p *UserEventProcessor) Handle(ctx context.Context, msg port.Message) error {
	fmt.Println("msg.Topic", msg.Topic)
	h, ok := p.handlers[msg.Topic]
	if !ok {
		return nil
	}

	var rawPayload json.RawMessage = msg.Value

	return h.Handle(ctx, rawPayload)
}
