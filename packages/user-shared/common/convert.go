package common

import (
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

func ToStringSlice[T ~string](items []T) []string {
	result := make([]string, len(items))
	for i, v := range items {
		result[i] = string(v)
	}

	return result
}

func ToEnumSlice[T ~string](items []string) []T {
	result := make([]T, len(items))
	for i, v := range items {
		result[i] = T(v)
	}

	return result
}
