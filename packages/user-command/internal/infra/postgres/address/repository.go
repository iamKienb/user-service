package address

import (
	"context"
	"user-command-module/db/repository"
	domain_address "user-command-module/internal/domain/address"

	pgx "github.com/iamKienb/go-core/postgres"
)

type userAddressRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) domain_address.Repository {
	return &userAddressRepository{
		queries: repository.New(service.GetPool()),
	}
}

func (r *userAddressRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}
