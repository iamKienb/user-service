package profile

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type UserProfileCreatedEvent struct {
	UserID    shared.UserID
	FullName  string
	Gender    GenderEnum
	CreatedAt time.Time
}

func (e UserProfileCreatedEvent) EventName() string {
	return "user-service.user.profile.created"
}

func (e UserProfileCreatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"user_id":    e.UserID.String(),
		"full_name":  e.FullName,
		"gender":     string(e.Gender),
		"created_at": e.CreatedAt,
	}
}
