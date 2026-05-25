package login

import (
	"context"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/auth"

	pgx "github.com/iamKienb/go-core/postgres"
)

type loginRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) auth.Repository {
	return &loginRepository{
		queries: repository.New(service.GetPool()),
	}
}

func (r *loginRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}

	return r.queries
}
