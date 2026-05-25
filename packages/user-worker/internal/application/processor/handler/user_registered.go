package handler

import (
	"context"
	"encoding/json"
	"user-shared-module/events"
	"user-worker-module/internal/application/port"
)

type UserRegisterHandler struct {
	repo  port.ESRepository
	alias string
}

func NewUserRegisterHandler(repo port.ESRepository, alias string) *UserRegisterHandler {
	return &UserRegisterHandler{repo: repo, alias: alias}
}

func (h *UserRegisterHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.UserRegistered
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	doc := map[string]any{
		"id":         payload.UserID,
		"email":      payload.Email,
		"status":     payload.Status,
		"roles":      payload.Roles,
		"created_at": payload.CreatedAt,
	}

	return h.repo.SyncData(ctx, h.alias, payload.UserID, doc)
}
