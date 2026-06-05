package handler

import (
	"context"
	"encoding/json"
	"user-shared-module/events"
	"user-worker-module/internal/application/port"
)

type UserAddressAddedHandler struct {
	repo  port.ESRepository
	alias string
}

func NewUserAddressAddedHandler(repo port.ESRepository, alias string) *UserAddressAddedHandler {
	return &UserAddressAddedHandler{repo: repo, alias: alias}
}

func (h *UserAddressAddedHandler) Handle(ctx context.Context, raw json.RawMessage) error {
	var payload events.UserAddressAdded
	if err := json.Unmarshal(raw, &payload); err != nil {
		return err
	}

	fullAddress := payload.AddressLine + payload.WardName + payload.ProvinceName

	doc := map[string]any{
		"address": map[string]any{
			"id":      payload.UserAddressID,
			"user_id": payload.UserID,

			"full_address": fullAddress,
			"country": map[string]any{
				"country_id": payload.CountryID,
				"name":       payload.CountryName,
			},
			"province": map[string]any{
				"id":   payload.ProvinceID,
				"name": payload.ProvinceName,
			},
			"ward": map[string]any{
				"id":   payload.WardID,
				"name": payload.WardName,
			},

			"address_line":  payload.AddressLine,
			"receiver_name": payload.ReceiverName,
			"phone_number":  payload.PhoneNumber,
			"label":         payload.Label,
			"is_default":    payload.IsDefault,
			"created_at":    payload.CreatedAt,
		},
	}

	return h.repo.SyncData(ctx, h.alias, payload.UserID, doc)
}
