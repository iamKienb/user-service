package events

import "time"

const TopicUserRegistered = "user-service.user.registered"

type UserRegistered struct {
	UserID    string    `json:"user_id"`
	Email     string    `json:"email"`
	Status    string    `json:"status"`
	FullName  string    `json:"full_name"`
	Gender    string    `json:"gender"`
	Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u UserRegistered) EventName() string {
	return TopicUserRegistered
}
