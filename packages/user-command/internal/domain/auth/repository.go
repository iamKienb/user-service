package auth

import "context"

type QueryRepository interface {
	FindLoginStatByUserID(ctx context.Context, userID string) (*LoginStat, error)
}

type CommandRepository interface {
	SaveLoginStat(ctx context.Context, stat *LoginStat) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
