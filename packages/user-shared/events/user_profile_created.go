package events

import "time"

const TopicUserProfileCreated = "user-service.user.profile.created"

type UserProfileCreated struct {
	UserID    string    `json:"user_id"`
	FullName  string    `json:"full_name"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
