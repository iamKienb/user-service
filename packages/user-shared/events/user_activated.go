package events

import "time"

const TopicUserActivated = "user-service.user.activated"

type UserActivated struct {
	UserID    string    `json:"user_id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (p UserActivated) EventName() string {
	return TopicUserActivated
}
