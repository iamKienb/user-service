package account

import (
	"github.com/google/uuid"
)

type UserID struct {
	Value uuid.UUID
}

func (id UserID) String() string {
	return id.Value.String()
}

func NewID() UserID {
	return UserID{Value: uuid.Must(uuid.NewV7())}
}
