package auth

import (
	"time"

	"user-command-module/internal/domain/shared"
)

const (
	MaxFailedAttempts  = 5
	LockoutDuration    = 30 * time.Minute
	DefaultFailedCount = 0
)

type LoginAttempt struct {
	UserID       shared.UserID
	FailedCount  int
	LastFailedAt *time.Time
	LockUntil    *time.Time
	UpdatedAt    time.Time
}

func NewAttempt(userID shared.UserID) *LoginAttempt {
	now := time.Now().UTC()

	return &LoginAttempt{
		UserID:       userID,
		FailedCount:  DefaultFailedCount,
		LastFailedAt: nil,
		LockUntil:    nil,
		UpdatedAt:    now,
	}
}

func (a *LoginAttempt) IsLocked() bool {
	if a.LockUntil == nil {
		return false
	}

	return time.Now().UTC().Before(*a.LockUntil)
}

func (a *LoginAttempt) EnsureCanAttemptLogin() error {
	if a.IsLocked() {
		return ErrAccountLocked
	}

	return nil
}

func (a *LoginAttempt) RecordFailure() {
	a.FailedCount++
	now := time.Now().UTC()
	a.LastFailedAt = &now
	a.UpdatedAt = now
	if a.FailedCount >= MaxFailedAttempts {
		lockUntil := now.Add(LockoutDuration)
		a.LockUntil = &lockUntil
	}
}

func (a *LoginAttempt) RecordSuccess() {
	now := time.Now().UTC()
	a.FailedCount = 0
	a.LastFailedAt = nil
	a.LockUntil = nil
	a.UpdatedAt = now
}
