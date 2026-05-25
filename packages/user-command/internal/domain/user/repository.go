package user

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	FindUserByEmail(ctx context.Context, email string) (*User, error)
	FindUserByID(ctx context.Context, userID shared.UserID) (*User, error)
}

type CommandRepository interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *User) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
