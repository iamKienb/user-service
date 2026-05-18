package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopMember struct {
	ShopID      shared.ShopID
	MemberID    shared.UserID
	JoinedAt    time.Time
	AddedBy     shared.UserID
	EventEntity shared.EventEntity
}

type ShopMemberRole struct {
	ShopID    shared.ShopID
	MemberID  shared.UserID
	RoleID    shared.RoleID
	UpdatedBy shared.UserID
}

type MemberPermission struct {
	ShopID  shared.ShopID
	RoleIDs []shared.RoleID
}
