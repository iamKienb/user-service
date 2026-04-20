package user

import (
	"context"
	"shopify-user-command-module/db/repository"
	"shopify-user-command-module/internal/domain/identity"

	postgresx "github.com/iamKienb/shopify-go-platform/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db      *pgxpool.Pool
	queries *repository.Queries
}

func NewUserRepository(db *pgxpool.Pool) identity.Repository {
	return &userRepository{
		db:      db,
		queries: repository.New(db),
	}
}

func (r *userRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := postgresx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}
