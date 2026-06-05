package login

import (
	"time"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toDomainLoginAttempt(row repository.LoginAttempt) *auth.LoginAttempt {
	var lockUntil *time.Time
	if row.LockUntil.Valid {
		lockUntil = &row.LockUntil.Time
	}

	var lastFailedAt *time.Time
	if row.LastFailedAt.Valid {
		lastFailedAt = &row.LastFailedAt.Time
	}

	return &auth.LoginAttempt{
		UserID:       shared.UserID(row.UserID.Bytes),
		FailedCount:  int(row.FailedCount),
		LastFailedAt: lastFailedAt,
		LockUntil:    lockUntil,
		UpdatedAt:    row.UpdatedAt.Time,
	}
}

func toInfraLoginAttempt(s *auth.LoginAttempt) repository.SaveLoginAttemptParams {
	return repository.SaveLoginAttemptParams{
		UserID:       conv.UUID(s.UserID),
		FailedCount:  int32(s.FailedCount),
		LockUntil:    conv.TimeStampZ(s.LockUntil),
		LastFailedAt: conv.TimeStampZ(s.LastFailedAt),
		UpdatedAt:    conv.TimeStampZ(&s.UpdatedAt),
	}
}
