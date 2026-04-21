package account

import "time"

type Aggregate struct {
	User       User
	Credential *Credential
	Profile    *Profile
}

func NewAggregate(p NewAggregateParams) *Aggregate {
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
		PasswordVersion:   DefaultPasswordVersion,
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

	return &Aggregate{
		User:       user,
		Credential: &credential,
		Profile:    &profile,
	}
}

func LoadAggregate(user User, credential *Credential, profile *Profile) *Aggregate {
	return &Aggregate{
		User:       user,
		Credential: credential,
		Profile:    profile,
	}
}

func (a *Aggregate) Activate() {
	a.User.Activate()
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

func (a *Aggregate) EnsureCanLogin() error {
	if err := a.EnsureCredential(); err != nil {
		return err
	}
	if !a.User.IsActive() {
		return ErrUserNotActive
	}

	return nil
}

func (a *Aggregate) ActivateIfNeeded() bool {
	if a.User.IsActive() {
		return false
	}

	a.User.Activate()
	return true
}
