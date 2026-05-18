package account

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	LoadAggByID(ctx context.Context, userID shared.UserID) (*Aggregate, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserID(ctx context.Context, userID shared.UserID) (*User, error)
	LoadAggByEmail(ctx context.Context, email string) (*Aggregate, error)
}

type CommandRepository interface {
	SaveAggregate(ctx context.Context, agg *Aggregate) error
	UpdateUser(ctx context.Context, user *User) error
	SaveAddress(ctx context.Context, addr *UserAddress) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
