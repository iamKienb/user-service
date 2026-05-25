package handler

import (
	"context"
	"encoding/json"
	"user-shared-module/events"
	"user-worker-module/internal/application/port"
)

type UserProfileCreatedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewUserProfileCreatedHandler(repo port.ESRepository, alias string) *UserProfileCreatedHandler {
	return &UserProfileCreatedHandler{repo: repo, alias: alias}
}

func (h *UserProfileCreatedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.UserProfileCreated
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	doc := map[string]any{
		"profile": map[string]any{
			"user_id":    payload.UserID,
			"full_name":  payload.FullName,
			"gender":     payload.Gender,
			"created_at": payload.CreatedAt,
		},
	}

	return h.repo.SyncData(ctx, h.alias, payload.UserID, doc)
}
