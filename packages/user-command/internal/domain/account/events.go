package account

import (
	"time"
	"user-command-module/internal/domain/shared"
	"user-shared-module/common"
)

type UserRegisteredEvent struct {
	UserID    shared.UserID
	Email     string
	Status    StatusEnum
	Roles     []RoleEnum
	FullName  string
	Gender    GenderEnum
	CreatedAt time.Time
}

func (e UserRegisteredEvent) EventName() string {
	return "user-service.user.registered"
}

func (e UserRegisteredEvent) IntegrationPayload() map[string]interface{} {

	return map[string]interface{}{
		"user_id":    e.UserID,
		"email":      e.Email,
		"status":     e.Status,
		"roles":      common.ToStringSlice(e.Roles),
		"full_name":  e.FullName,
		"gender":     string(e.Gender),
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
		"user_id":    e.UserID,
		"status":     e.Status,
		"created_at": e.UpdatedAt,
	}
}
