package auth

import (
	"context"

	"shopify-user-command-module/db/repository"
	"shopify-user-command-module/internal/domain/auth"

	postgresx "github.com/iamKienb/shopify-go-platform/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type authRepository struct {
	queries *repository.Queries
}

func NewRepository(db *pgxpool.Pool) auth.Repository {
	return &authRepository{
		queries: repository.New(db),
	}
}

func (r *authRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := postgresx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}

	return r.queries
}
