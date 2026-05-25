package user

import (
	"context"
	"errors"
	"user-command-module/db/repository"
	domain_user "user-command-module/internal/domain/user"

	pgx "github.com/iamKienb/go-core/postgres"
	"github.com/jackc/pgx/v5/pgconn"
)

type userRepository struct {
	queries *repository.Queries
}

func NewRepository(service pgx.PGXService) domain_user.Repository {
	return &userRepository{
		queries: repository.New(service.GetPool()),
	}
}

func (r *userRepository) getQuerier(ctx context.Context) *repository.Queries {
	if tx := pgx.ExtractTx(ctx); tx != nil {
		return r.queries.WithTx(tx)
	}
	return r.queries
}

func (r *userRepository) IsDuplicateEmail(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" && pgErr.ConstraintName == "uq_users_email"
	}

	return false
}
