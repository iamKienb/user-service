package auth

import (
	"time"

	"shopify-user-command-module/internal/domain/account"
)

const (
	MaxFailedAttempts  = 5
	LockoutDuration    = 30 * time.Minute
	DefaultFailedCount = 0
)

type LoginStat struct {
	UserID       account.UserID
	FailedCount  int
	LastFailedAt *time.Time
	LockUntil    *time.Time
	UpdatedAt    time.Time
}

func NewLoginStat(userID account.UserID) *LoginStat {
	now := time.Now().UTC()

	return &LoginStat{
		UserID:       userID,
		FailedCount:  DefaultFailedCount,
		LastFailedAt: nil,
		LockUntil:    nil,
		UpdatedAt:    now,
	}
}

func (s *LoginStat) IsLocked() bool {
	if s.LockUntil == nil {
		return false
	}

	return time.Now().UTC().Before(*s.LockUntil)
}

func (s *LoginStat) EnsureCanAttemptLogin() error {
	if s.IsLocked() {
		return ErrAccountLocked
	}

	return nil
}

func (s *LoginStat) RecordFailure() {
	s.FailedCount++
	now := time.Now().UTC()
	s.LastFailedAt = &now
	s.UpdatedAt = now
	if s.FailedCount >= MaxFailedAttempts {
		lockUntil := now.Add(LockoutDuration)
		s.LockUntil = &lockUntil
	}
}

func (s *LoginStat) RecordSuccess() {
	now := time.Now().UTC()
	s.FailedCount = 0
	s.LastFailedAt = nil
	s.LockUntil = nil
	s.UpdatedAt = now
}
