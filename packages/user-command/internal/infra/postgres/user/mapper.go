package user

import (
	"time"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/shared"
	domain_user "user-command-module/internal/domain/user"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toInfraUser(u *domain_user.User) repository.CreateUserParams {
	return repository.CreateUserParams{
		ID:              conv.UUID(u.ID),
		Email:           u.Email,
		EmailVerifiedAt: conv.TimeStampZ(u.EmailVerifiedAt),
		Status:          string(u.Status),
		Roles:           shared.Strings(u.Roles),
		CreatedAt:       conv.TimeStampZ(&u.CreatedAt),
		UpdatedAt:       conv.TimeStampZ(u.UpdatedAt),
		DeletedAt:       conv.TimeStampZ(u.DeletedAt),
	}
}

func toDomainUser(userRow repository.User, credentialRow *repository.UserCredential) *domain_user.User {
	var emailVerifiedAt *time.Time
	if userRow.EmailVerifiedAt.Valid {
		emailVerifiedAt = &userRow.EmailVerifiedAt.Time
	}

	var deletedAt *time.Time
	if userRow.DeletedAt.Valid {
		deletedAt = &userRow.DeletedAt.Time
	}

	credential := domain_user.UserCredential{
		UserID:            credentialRow.UserID.Bytes,
		PasswordHash:      credentialRow.PasswordHash,
		PasswordVersion:   int(credentialRow.PasswordVersion),
		PasswordUpdatedAt: credentialRow.PasswordUpdatedAt.Time,
	}

	return &domain_user.User{
		ID:              shared.UserID(userRow.ID.Bytes),
		Email:           userRow.Email,
		Status:          domain_user.StatusEnum(userRow.Status),
		EmailVerifiedAt: emailVerifiedAt,
		Roles:           shared.FromStrings[domain_user.RoleEnum](userRow.Roles),
		Credential:      credential,
		CreatedAt:       userRow.CreatedAt.Time,
		UpdatedAt:       &userRow.UpdatedAt.Time,
		DeletedAt:       deletedAt,
	}
}

func toUpdateUserInfra(u *domain_user.User) repository.UpdateUserParams {
	now := time.Now().UTC()
	return repository.UpdateUserParams{
		ID:              conv.UUID(u.ID),
		Email:           u.Email,
		Status:          string(u.Status),
		EmailVerifiedAt: conv.TimeStampZ(u.EmailVerifiedAt),
		Roles:           shared.Strings(u.Roles),
		UpdatedAt:       conv.TimeStampZ(&now),
		DeletedAt:       conv.TimeStampZ(u.DeletedAt),
	}
}

func toDomainCredential(row repository.UserCredential) *domain_user.UserCredential {
	return &domain_user.UserCredential{
		UserID:            shared.UserID(row.UserID.Bytes),
		PasswordHash:      row.PasswordHash,
		PasswordVersion:   int(row.PasswordVersion),
		PasswordUpdatedAt: row.PasswordUpdatedAt.Time,
	}
}

func toInfraCredential(c *domain_user.UserCredential) repository.CreateUserCredentialParams {
	return repository.CreateUserCredentialParams{
		UserID:            conv.UUID(c.UserID),
		PasswordHash:      c.PasswordHash,
		PasswordVersion:   int32(c.PasswordVersion),
		PasswordUpdatedAt: conv.TimeStampZ(&c.PasswordUpdatedAt),
	}
}
