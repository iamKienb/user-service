package account

import "time"

type UserRole string

const (
	RoleCustomer   UserRole = "CUSTOMER"
	RoleShopStaff  UserRole = "SHOP_STAFF"
	RoleShopOwner  UserRole = "SHOP_OWNER"
	RoleSuperAdmin UserRole = "SUPER_ADMIN"
)

type UserStatus string

const (
	StatusPending UserStatus = "PENDING"
	StatusActive  UserStatus = "ACTIVE"
	StatusBanned  UserStatus = "BANNED"
	StatusDeleted UserStatus = "DELETED"
)

type User struct {
	ID              UserID
	Email           string
	Status          UserStatus
	EmailVerifiedAt *time.Time
	Roles           []UserRole
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}

func (u *User) IsActive() bool {
	return u.Status == StatusActive
}

func (u *User) Activate() {
	u.Status = StatusActive
	now := time.Now().UTC()
	u.UpdatedAt = now
	if u.EmailVerifiedAt == nil {
		u.EmailVerifiedAt = &now
	}
}
