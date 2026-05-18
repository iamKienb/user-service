package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopStatus string

const (
	StatusPending  ShopStatus = "PENDING"
	StatusActive   ShopStatus = "ACTIVE"
	StatusInActive ShopStatus = "INACTIVE"
	StatusDeleted  ShopStatus = "DELETED"
)

type Shop struct {
	ID      shared.ShopID
	OwnerID shared.UserID
	Name    string
	Slug    string
	Status  ShopStatus

	CreatedBy shared.UserID
	UpdatedBy *shared.UserID

	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
	EventEntity shared.EventEntity
}

func (s *Shop) MatchesSlug(slug string) bool {
	return s.Slug == slug
}

func (s *Shop) IsActive() bool {
	return s.Status == StatusActive
}
