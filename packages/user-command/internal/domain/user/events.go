package user

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type UserRegisteredEvent struct {
	UserID    shared.UserID
	Email     string
	Status    StatusEnum
	Roles     []RoleEnum
	CreatedAt time.Time
}

func (e UserRegisteredEvent) EventName() string {
	return "user-service.user.registered"
}

func (e UserRegisteredEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"user_id":    e.UserID.String(),
		"email":      e.Email,
		"status":     string(e.Status),
		"roles":      shared.Strings(e.Roles),
		"created_at": e.CreatedAt,
	}
}

type UserActivatedEvent struct {
	UserID    shared.UserID
	Status    StatusEnum
	UpdatedAt *time.Time
}

func (e UserActivatedEvent) EventName() string {
	return "user-service.user.activated"
}

func (e UserActivatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"user_id":    e.UserID.String(),
		"status":     e.Status,
		"updated_at": e.UpdatedAt,
	}
}
