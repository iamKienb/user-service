package auth

import (
	"time"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shared"
	"user-shared-module/common"
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
		UserID:       shared.UserID(row.UserID.Bytes),
		FailedCount:  int(row.FailedCount),
		LastFailedAt: lastFailedAt,
		LockUntil:    lockUntil,
		UpdatedAt:    row.UpdatedAt.Time,
	}
}

func toInfraLoginStat(s *auth.LoginStat) repository.SaveLoginStatsParams {
	return repository.SaveLoginStatsParams{
		UserID:       common.ToPgUUID(s.UserID),
		FailedCount:  int32(s.FailedCount),
		LockUntil:    common.ToPgTimeStampZ(s.LockUntil),
		LastFailedAt: common.ToPgTimeStampZ(s.LastFailedAt),
		UpdatedAt:    common.ToPgTimeStampZ(&s.UpdatedAt),
	}
}
