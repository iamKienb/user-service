package auth

import (
	"context"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/auth"

	pgx "github.com/iamKienb/go-core/postgres"
)

type authRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) auth.Repository {
	return &authRepository{
		queries: repository.New(service.GetPool()),
	}
}

func (r *authRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}

	return r.queries
}
