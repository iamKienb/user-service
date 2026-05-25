package user

import (
	"time"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/shared"
	domain_user "user-command-module/internal/domain/user"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toInfraUser(u *domain_user.User) repository.SaveUserParams {
	return repository.SaveUserParams{
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

func toDomainUser(row repository.User) *domain_user.User {
	var emailVerifiedAt *time.Time
	if row.EmailVerifiedAt.Valid {
		emailVerifiedAt = &row.EmailVerifiedAt.Time
	}

	var deletedAt *time.Time
	if row.DeletedAt.Valid {
		deletedAt = &row.DeletedAt.Time
	}

	return &domain_user.User{
		ID:              shared.UserID(row.ID.Bytes),
		Email:           row.Email,
		Status:          domain_user.StatusEnum(row.Status),
		EmailVerifiedAt: emailVerifiedAt,
		Roles:           shared.FromStrings[domain_user.RoleEnum](row.Roles),
		CreatedAt:       row.CreatedAt.Time,
		UpdatedAt:       &row.UpdatedAt.Time,
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

func toInfraCredential(c *domain_user.UserCredential) repository.SaveUserCredentialParams {
	return repository.SaveUserCredentialParams{
		UserID:            conv.UUID(c.UserID),
		PasswordHash:      c.PasswordHash,
		PasswordVersion:   int32(c.PasswordVersion),
		PasswordUpdatedAt: conv.TimeStampZ(&c.PasswordUpdatedAt),
	}
}
