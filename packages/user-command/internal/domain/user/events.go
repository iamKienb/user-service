package user

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type UserRegisteredEvent struct {
	UserID    string
	Email     string
	Status    StatusEnum
	Roles     []RoleEnum
	FullName  string
	Gender    string
	CreatedAt time.Time
}

func (e UserRegisteredEvent) EventName() string {
	return "user-service.user.registered"
}

func (e UserRegisteredEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"user_id": e.UserID,
		"email":   e.Email,
		"status":  string(e.Status),
		"roles":   shared.Strings(e.Roles),
		"profile": map[string]interface{}{
			"user_id":    e.UserID,
			"full_name":  e.FullName,
			"gender":     e.Gender,
			"created_at": e.CreatedAt,
		},
		"created_at": e.CreatedAt,
	}
}

type UserActivatedEvent struct {
	UserID    string
	Status    StatusEnum
	UpdatedAt *time.Time
}

func (e UserActivatedEvent) EventName() string {
	return "user-service.user.activated"
}

func (e UserActivatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"user_id":    e.UserID,
		"status":     e.Status,
		"updated_at": e.UpdatedAt,
	}
}
