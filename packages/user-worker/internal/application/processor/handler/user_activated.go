package handler

import (
	"context"
	"encoding/json"
	"user-shared-module/events"
	"user-worker-module/internal/application/port"
)

type UserActivatedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewUserActivatedHandler(repo port.ESRepository, alias string) *UserActivatedHandler {
	return &UserActivatedHandler{repo: repo, alias: alias}
}

func (h *UserActivatedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.UserActivated
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	doc := map[string]interface{}{
		"status":     payload.Status,
		"updated_at": payload.UpdatedAt,
	}

	return h.repo.SyncData(ctx, h.alias, payload.UserID, doc)
}
