package account

import "context"

type QueryRepository interface {
	FindAggregateByID(ctx context.Context, id string) (*Aggregate, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserID(ctx context.Context, id string) (*User, error)
	FindForAuthentication(ctx context.Context, email string) (*Aggregate, error)
}

type CommandRepository interface {
	SaveAggregate(ctx context.Context, agg *Aggregate) error
	UpdateUser(ctx context.Context, user *User) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
