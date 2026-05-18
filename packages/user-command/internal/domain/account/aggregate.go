package account

import (
	"time"
	"user-command-module/internal/domain/shared"
)

const AggregateTypeUser = "USER"

type Aggregate struct {
	User       User
	Credential *UserCredential
	Profile    *UserProfile
	Events     []shared.DomainEvent
}

func NewAggregate(p AggregateParams) *Aggregate {
	now := time.Now().UTC()

	userID := shared.NewID[shared.UserID]()
	roles := []RoleEnum{RoleCustomer}

	user := User{
		ID:              userID,
		Email:           p.Email,
		EmailVerifiedAt: nil,
		Status:          StatusPending,
		Roles:           roles,
		CreatedAt:       now,
		UpdatedAt:       nil,
		DeletedAt:       nil,
	}

	credential := UserCredential{
		UserID:            userID,
		PasswordHash:      p.PasswordHash,
		PasswordVersion:   DefaultPasswordVersion,
		PasswordUpdatedAt: now,
	}

	profile := UserProfile{
		UserID:      userID,
		FullName:    p.FullName,
		Gender:      p.Gender,
		PhoneNumber: nil,
		AvatarURL:   nil,
		DateOfBirth: nil,
		CreatedAt:   now,
		UpdatedAt:   nil,
	}

	aggregate := &Aggregate{
		User:       user,
		Credential: &credential,
		Profile:    &profile,
	}

	aggregate.User.EventEntity.AddEvent(UserRegisteredEvent{
		UserID:    userID,
		Email:     aggregate.User.Email,
		Status:    aggregate.User.Status,
		Roles:     aggregate.User.Roles,
		FullName:  aggregate.Profile.FullName,
		Gender:    aggregate.Profile.Gender,
		CreatedAt: aggregate.User.CreatedAt,
	})

	return aggregate
}

func LoadAggregate(user User, credential *UserCredential, profile *UserProfile) *Aggregate {
	return &Aggregate{
		User:       user,
		Credential: credential,
		Profile:    profile,
	}
}

func (a *Aggregate) FlushEvents() []shared.DomainEvent {
	var allEvents []shared.DomainEvent
	if len(a.Events) > 0 {
		allEvents = append(allEvents, a.Events...)
		a.Events = nil
	}

	allEvents = append(allEvents, a.User.EventEntity.Flush()...)
	return allEvents
}

func (a *Aggregate) EnsureCredential() error {
	if a == nil {
		return ErrUserNotFound
	}
	if a.Credential == nil {
		return ErrCredentialNotFound
	}

	return nil
}

func (a *Aggregate) CheckActiveIfLogin() error {
	if err := a.EnsureCredential(); err != nil {
		return err
	}
	if !a.User.IsActive() {
		return ErrUserNotActive
	}

	return nil
}

func (a *Aggregate) ActivateIfVerified() bool {
	if a.User.IsActive() {
		return false
	}

	a.User.Activate()

	a.User.EventEntity.AddEvent(UserActivatedEvent{
		UserID:    a.User.ID,
		Status:    a.User.Status,
		UpdatedAt: a.User.UpdatedAt,
	})

	return true
}
