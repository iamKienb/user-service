package processor

import (
	"context"
	"encoding/json"
	"user-shared-module/alias"
	"user-shared-module/events"
	"user-worker-module/internal/application/port"
	"user-worker-module/internal/application/processor/handler"
)

type UserEventProcessor struct {
	handlers map[string]port.EventHandler
}

func NewUserEventProcessor(repo port.ESRepository) port.EventProcessor {
	userAlias := alias.UserAlias

	p := &UserEventProcessor{
		handlers: make(map[string]port.EventHandler),
	}

	p.handlers[events.TopicUserRegistered] = handler.NewUserRegisterHandler(repo, userAlias)
	p.handlers[events.TopicUserActivated] = handler.NewUserActivatedHandler(repo, userAlias)
	p.handlers[events.TopicUserProfileCreated] = handler.NewUserProfileCreatedHandler(repo, userAlias)
	p.handlers[events.TopicUserAddressAdded] = handler.NewUserAddressAddedHandler(repo, userAlias)

	return p
}

func (p *UserEventProcessor) Handle(ctx context.Context, msg port.Message) error {
	h, ok := p.handlers[msg.Topic]
	if !ok {
		return nil
	}

	var rawPayload json.RawMessage = msg.Value

	return h.Handle(ctx, rawPayload)
}
