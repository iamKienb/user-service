package auth

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	FindLoginStatByUserID(ctx context.Context, userID shared.UserID) (*LoginStat, error)
}

type CommandRepository interface {
	SaveLoginStat(ctx context.Context, stat *LoginStat) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
