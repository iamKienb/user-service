package user

type DomainEvent interface {
	EventName() string
}

type UserRegistered struct {
	UserID string
	Email  string
}

func (e UserRegistered) EventName() string {
	return "user.registered"
}
