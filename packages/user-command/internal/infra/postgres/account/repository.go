package account

import (
	"context"
	"errors"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/account"

	pgx "github.com/iamKienb/go-core/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type accountRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) account.Repository {
	return &accountRepository{
		queries: repository.New(service.GetPool()),
	}
}

func (r *accountRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}

func (r *accountRepository) IsDuplicateEmail(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" && pgErr.ConstraintName == "uq_users_email"
	}

	return false
}
