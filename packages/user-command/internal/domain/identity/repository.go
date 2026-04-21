package identity

import (
	"context"
)

type QueryRepo interface {
	FindAggregateByID(ctx context.Context, id string) (*IdentityAggregate, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserID(ctx context.Context, id string) (*User, error)
	FindForLogin(ctx context.Context, email string) (*IdentityAggregate, error)
	FindLoginStatByID(ctx context.Context, id string) (*LoginStat, error)
}

type CommandRepo interface {
	SaveAggregate(ctx context.Context, agg *IdentityAggregate) error
	SaveLoginStat(ctx context.Context, params *LoginStat) error
	UpdateUser(ctx context.Context, params *User) error
}

type Repository interface {
	QueryRepo
	CommandRepo
}
