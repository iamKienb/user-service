package identity

import "time"

const (
	MaxFailedAttempts  = 5
	LockoutDuration    = 30 * time.Minute
	DefaultFailedCount = 0
)

type LoginStat struct {
	UserID       UserID
	FailedCount  int
	LastFailedAt *time.Time
	LockUntil    *time.Time
	UpdatedAt    time.Time
}

func NewLoginStat(userID UserID) *LoginStat {
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
		return true
	}

	return s.LockUntil.After(time.Now())
}

func (s *LoginStat) RecordFailure() {
	s.FailedCount++
	now := time.Now()
	if s.FailedCount >= MaxFailedAttempts {
		lockUntil := now.Add(LockoutDuration)
		s.LockUntil = &lockUntil
	}
}

func (s *LoginStat) RecordSuccess() {
	s.FailedCount = 0
	s.LockUntil = nil
}
