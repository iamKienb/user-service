package shop

import (
	"context"
	"errors"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/shop"

	pgx "github.com/iamKienb/go-core/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type shopRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) shop.Repository {
	return &shopRepository{
		queries: repository.New(service.GetPool()),
	}
}
func (r *shopRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}

func (r *shopRepository) IsDuplicateSlug(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" && pgErr.ConstraintName == "uq_shops_slug"
	}

	return false
}
