package profile

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	FindProfileByID(ctx context.Context, userID shared.UserID) (*Profile, error)
}

type CommandRepository interface {
	CreateProfile(ctx context.Context, profile *Profile) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
