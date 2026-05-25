package profile

import (
	"context"
	"user-command-module/db/repository"
	domain_profile "user-command-module/internal/domain/profile"

	pgx "github.com/iamKienb/go-core/postgres"
)

type profileRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) domain_profile.Repository {
	return &profileRepository{
		queries: repository.New(service.GetPool()),
	}
}

func (r *profileRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}
