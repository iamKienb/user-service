package user

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type RoleEnum string

const (
	RoleCustomer   RoleEnum = "CUSTOMER"
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
	Credential      UserCredential

	CreatedAt time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time

	shared.EventEntity
}

func NewUser(params NewUserParams) *User {
	now := time.Now().UTC()
	userID := shared.NewID[shared.UserID]()
	roles := []RoleEnum{RoleCustomer}

	credential := UserCredential{
		UserID:            userID,
		PasswordHash:      params.PasswordHash,
		PasswordVersion:   DefaultPasswordVersion,
		PasswordUpdatedAt: now,
	}

	user := &User{
		ID:              userID,
		Email:           params.Email,
		EmailVerifiedAt: nil,
		Status:          StatusPending,
		Roles:           roles,
		Credential:      credential,
		CreatedAt:       now,
		UpdatedAt:       nil,
		DeletedAt:       nil,
	}

	user.AddEvent(UserRegisteredEvent{
		UserID:    userID,
		Email:     user.Email,
		Status:    user.Status,
		Roles:     user.Roles,
		CreatedAt: user.CreatedAt,
	})

	return user
}

func (a *User) FlushEvents() []shared.DomainEvent {
	var domainEvents []shared.DomainEvent

	domainEvents = append(domainEvents, a.EventEntity.Flush()...)
	a.EventEntity.ClearEvent()
	return domainEvents
}

func (u *User) IsActive() bool {
	return u.Status == StatusActive
}

func (u *User) EnsureActiveForLogin() error {
	if !u.IsActive() {
		return ErrUserNotActive
	}

	return nil
}

func (u *User) ActivateIfVerified() bool {
	if u.IsActive() {
		return false
	}

	u.Status = StatusActive
	now := time.Now().UTC()
	u.UpdatedAt = &now

	if u.EmailVerifiedAt == nil {
		u.EmailVerifiedAt = &now
	}

	u.EventEntity.AddEvent(UserActivatedEvent{
		UserID:    u.ID,
		Status:    u.Status,
		UpdatedAt: u.UpdatedAt,
	})

	return true
}

func (u User) Type() string {
	return "USER"
}
