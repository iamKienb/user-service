package common

import (
	"shopify-user-command-module/internal/domain/account"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(id [16]byte) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

func ToPgText(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{Valid: false}
	}

	return pgtype.Text{String: *s, Valid: true}
}

func ToPgDate(d *time.Time) pgtype.Date {
	if d == nil {
		return pgtype.Date{Valid: false}
	}

	return pgtype.Date{Time: *d, Valid: true}
}

func ToPgTimeStampZ(t *time.Time) pgtype.Timestamptz {
	if t == nil {
		return pgtype.Timestamptz{Valid: false}
	}

	return pgtype.Timestamptz{Time: *t, Valid: true}
}

func ToStringRoles(roles []account.UserRole) []string {
	result := make([]string, len(roles))
	for i, r := range roles {
		result[i] = string(r)
	}

	return result
}

func ToDomainRoles(roles []string) []account.UserRole {
	result := make([]account.UserRole, len(roles))
	for i, r := range roles {
		result[i] = account.UserRole(r)
	}

	return result
}
