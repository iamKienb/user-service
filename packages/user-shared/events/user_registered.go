package events

import "time"

const TopicUserRegistered = "user-service.user.registered"

type UserRegistered struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	Roles     []string  `json:"roles"`
	Profile   Profile   `json:"profile"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Profile struct {
	UserID    string    `json:"user_id"`
	FullName  string    `json:"full_name"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
}
