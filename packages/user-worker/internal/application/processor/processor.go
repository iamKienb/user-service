package processor

import (
	"context"
	"fmt"
	"time"
	"user-shared-module/alias"
	"user-shared-module/events"
	"user-worker-module/internal/application/port"
	"user-worker-module/internal/application/processor/handler"
)

const (
	idemKeyTTL = 24 * time.Hour
	key        = "user-worker:key:%s"
)

type UserEventProcessor struct {
	handlers map[string]port.EventHandler
	port.WorkerCache
}

func NewUserEventProcessor(repo port.ESRepository, workerCache port.WorkerCache) port.EventProcessor {
	userAlias := alias.UserAlias

	p := &UserEventProcessor{
		handlers:    make(map[string]port.EventHandler),
		WorkerCache: workerCache,
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

	idemKey := msg.IdempotencyKey()

	if idemKey != "" {
		key := fmt.Sprintf(key, idemKey)
		isNew, err := p.WorkerCache.SetNx(ctx, key, 1, idemKeyTTL)
		if err != nil {
			return err
		}

		if !isNew {
			return nil
		}
	}

	return h.Handle(ctx, msg.Value)
}
