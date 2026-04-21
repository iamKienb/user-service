package identity

import (
	"time"

	"github.com/google/uuid"
)

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

type UserID struct {
	Value uuid.UUID
}

func (id UserID) String() string {
	return id.Value.String()
}

func NewID() UserID {
	return UserID{uuid.Must(uuid.NewV7())}
}

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

func (u *User) HasRoles() []UserRole {
	return u.Roles
}

func (u *User) Activate() {
	u.Status = StatusActive
}

type Credential struct {
	UserID            UserID
	PasswordHash      string
	PasswordVersion   int
	PasswordUpdatedAt time.Time
}

type Profile struct {
	UserID      UserID
	FullName    string
	Gender      string
	PhoneNumber *string
	AvatarURL   *string
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
