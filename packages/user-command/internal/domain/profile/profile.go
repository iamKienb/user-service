package profile

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type GenderEnum string

const (
	GenderMale   GenderEnum = "MALE"
	GenderFemale GenderEnum = "FEMALE"
	GenderOther  GenderEnum = "OTHER"
)

type Profile struct {
	UserID      shared.UserID
	FullName    string
	Gender      GenderEnum
	PhoneNumber *string
	AvatarURL   *string
	DateOfBirth *time.Time
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	shared.EventEntity
}

func NewProfile(params NewProfileParams) *Profile {
	now := time.Now().UTC()

	profile := &Profile{
		UserID:      params.UserID,
		FullName:    params.FullName,
		Gender:      params.Gender,
		PhoneNumber: nil,
		AvatarURL:   nil,
		DateOfBirth: nil,
		CreatedAt:   now,
		UpdatedAt:   nil,
	}

	profile.AddEvent(UserProfileCreatedEvent{
		UserID:    profile.UserID,
		FullName:  profile.FullName,
		Gender:    profile.Gender,
		CreatedAt: profile.CreatedAt,
	})

	return profile
}

func (p *Profile) FlushEvents() []shared.DomainEvent {
	var domainEvents []shared.DomainEvent
	domainEvents = append(domainEvents, p.EventEntity.Flush()...)
	p.ClearEvent()

	return domainEvents
}

func (e GenderEnum) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOther:
		return true
	}

	return false
}

func (u Profile) Type() string {
	return "Profile"
}
