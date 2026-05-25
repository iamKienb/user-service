package auth

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	FindLoginAttemptByID(ctx context.Context, userID shared.UserID) (*LoginAttempt, error)
}

type CommandRepository interface {
	SaveLoginAttempt(ctx context.Context, attempt *LoginAttempt) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
