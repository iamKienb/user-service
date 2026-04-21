package common

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func IsDuplicateEmail(err error) bool {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		return pgErr.Code == "23505" && pgErr.ConstraintName == "uq_users_email"
	}

	return false
}
