package identity

import (
	"context"
)

type QueryRepo interface {
	FindAggregateByID(ctx context.Context, id string) (*IdentityAggregate, error)

	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUserID(ctx context.Context, id string) (*User, error)
	// FindForLogin(ctx context.Context, email string) (*User, *Credential, error)
}

type CommandRepo interface {
	Save(ctx context.Context, agg *IdentityAggregate) error

	UpdateCredential(ctx context.Context, params UpdateCredentialParams) error

	UpdateUser(ctx context.Context, params *User) error
}

type Repository interface {
	QueryRepo
	CommandRepo
}
