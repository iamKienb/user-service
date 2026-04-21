package auth

import (
	"time"

	"shopify-user-command-module/db/repository"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"
	"shopify-user-command-module/internal/infra/common"
)

func toDomainLoginStat(row repository.LoginStat) *auth.LoginStat {
	var lockUntil *time.Time
	if row.LockUntil.Valid {
		lockUntil = &row.LockUntil.Time
	}

	var lastFailedAt *time.Time
	if row.LastFailedAt.Valid {
		lastFailedAt = &row.LastFailedAt.Time
	}

	return &auth.LoginStat{
		UserID:       account.UserID{Value: row.UserID.Bytes},
		FailedCount:  int(row.FailedCount),
		LastFailedAt: lastFailedAt,
		LockUntil:    lockUntil,
		UpdatedAt:    row.UpdatedAt.Time,
	}
}

func toInfraLoginStat(s *auth.LoginStat) repository.SaveLoginStatsParams {
	return repository.SaveLoginStatsParams{
		UserID:       common.ToPgUUID(s.UserID.Value),
		FailedCount:  int32(s.FailedCount),
		LockUntil:    common.ToPgTimeStampZ(s.LockUntil),
		LastFailedAt: common.ToPgTimeStampZ(s.LastFailedAt),
		UpdatedAt:    common.ToPgTimeStampZ(&s.UpdatedAt),
	}
}
