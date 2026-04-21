package account

import (
	"context"

	"shopify-user-command-module/db/repository"
	"shopify-user-command-module/internal/domain/account"

	postgresx "github.com/iamKienb/shopify-go-platform/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type accountRepository struct {
	queries *repository.Queries
}

func NewRepository(db *pgxpool.Pool) account.Repository {
	return &accountRepository{
		queries: repository.New(db),
	}
}

func (r *accountRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := postgresx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}
