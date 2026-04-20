package identity

import (
	"shopify-user-command-module/internal/application/shared"
	"time"
)

type IdentityAggregate struct {
	User       User
	Credential *Credential
	Profile    *Profile
}

func NewAggregate(p NewAggregateParams) *IdentityAggregate {
	now := time.Now().UTC()
	userID := NewID()
	user := User{
		ID:              userID,
		Email:           p.Email,
		EmailVerifiedAt: nil,
		Status:          StatusPending,
		Roles:           []UserRole{RoleCustomer},
		CreatedAt:       now,
		UpdatedAt:       now,
		DeletedAt:       nil,
	}

	credential := Credential{
		UserID:            userID,
		PasswordHash:      p.PasswordHash,
		PasswordVersion:   shared.DefaultPasswordVersion,
		PasswordUpdatedAt: now,
	}

	profile := Profile{
		UserID:      userID,
		FullName:    p.FullName,
		Gender:      p.Gender,
		PhoneNumber: nil,
		AvatarURL:   nil,
		DateOfBirth: nil,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	return &IdentityAggregate{
		User:       user,
		Credential: &credential,
		Profile:    &profile,
	}
}

func LoadAggregate(user User, credential *Credential, profile *Profile) *IdentityAggregate {
	return &IdentityAggregate{
		User:       user,
		Credential: credential,
		Profile:    profile,
	}
}
