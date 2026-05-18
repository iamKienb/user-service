package account

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type RoleEnum string

const (
	RoleCustomer   RoleEnum = "CUSTOMER"
	RoleShopStaff  RoleEnum = "SHOP_STAFF"
	RoleShopOwner  RoleEnum = "SHOP_OWNER"
	RoleSuperAdmin RoleEnum = "SUPER_ADMIN"
)

type StatusEnum string

const (
	StatusPending StatusEnum = "PENDING"
	StatusActive  StatusEnum = "ACTIVE"
	StatusBanned  StatusEnum = "BANNED"
	StatusDeleted StatusEnum = "DELETED"
)

type User struct {
	ID              shared.UserID
	Email           string
	Status          StatusEnum
	EmailVerifiedAt *time.Time
	Roles           []RoleEnum
	CreatedAt       time.Time
	UpdatedAt       *time.Time
	DeletedAt       *time.Time
	EventEntity     shared.EventEntity
}

func (u *User) IsActive() bool {
	return u.Status == StatusActive
}

func (u *User) Activate() {
	u.Status = StatusActive
	now := time.Now().UTC()
	u.UpdatedAt = &now

	if u.EmailVerifiedAt == nil {
		u.EmailVerifiedAt = &now
	}
}
