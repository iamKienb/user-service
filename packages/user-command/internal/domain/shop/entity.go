package user

import "time"

type UserRole int

const (
	RoleCustomer UserRole = iota + 1
	RoleShopStaff
	RoleShopOwner
	RoleSuperAdmin
)

type User struct {
	id        UserID
	email     Email
	status    UserStatus
	roles     []UserRole
	createdAt time.Time
	UpdatedAt time.Time
	events    []DomainEvent
}

type UserStatus int

const (
	StatusPending UserStatus = iota + 1
	StatusActive
	StatusSuspended
	StatusDeleted
)

type UserCredential struct {
	userID            string
	password_hash     string
	IsLocked          bool
	LockUntil         *time.Time
	PasswordVersion   int
	PasswordUpdatedAt time.Time
}

type UserProfile struct {
	userID      string
	first_name  string
	last_name   string
	avatarURL   *string
	dateOfBirth *time.Time
	updatedAt   time.Time
}

type LoginStats struct {
	UserID            string
	FailedLoginCount  int
	LastLoginAt       *time.Time
	LastFailedLoginAt *time.Time
}

type Address struct {
	ID     string
	UserID string

	Label        string
	ReceiverName string
	PhoneNumber  string

	AddressLine string
	Ward        string
	District    string
	City        string
	Country     string

	IsDefault bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) ID() UserID {
	return UserID(u.id)
}

func (u *User) Email() Email {
	return Email(u.email)
}

func (u *User) Status() UserStatus {
	return UserStatus(u.status)
}

func (u *User) addEvent(e DomainEvent) {
	u.events = append(u.events, e)
}

func (u *User) PullEvents() []DomainEvent {
	events := u.events
	u.events = nil
	return events
}
