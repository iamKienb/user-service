package account

import (
	"time"

	"shopify-user-command-module/db/repository"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/infra/common"
)

func toDomainUser(row repository.User) *account.User {
	var emailVerifiedAt *time.Time
	if row.EmailVerifiedAt.Valid {
		emailVerifiedAt = &row.EmailVerifiedAt.Time
	}

	var deletedAt *time.Time
	if row.DeletedAt.Valid {
		deletedAt = &row.DeletedAt.Time
	}

	return &account.User{
		ID:              account.UserID{Value: row.ID.Bytes},
		Email:           row.Email,
		Status:          account.UserStatus(row.Status),
		EmailVerifiedAt: emailVerifiedAt,
		Roles:           common.ToDomainRoles(row.Roles),
		CreatedAt:       row.CreatedAt.Time,
		UpdatedAt:       row.UpdatedAt.Time,
		DeletedAt:       deletedAt,
	}
}

func toDomainCredential(row repository.Credential) *account.Credential {
	return &account.Credential{
		UserID:            account.UserID{Value: row.UserID.Bytes},
		PasswordHash:      row.PasswordHash,
		PasswordVersion:   int(row.PasswordVersion),
		PasswordUpdatedAt: row.PasswordUpdatedAt.Time,
	}
}

func toDomainProfile(row repository.Profile) *account.Profile {
	var phoneNumber *string
	if row.PhoneNumber.Valid {
		phoneNumber = &row.PhoneNumber.String
	}

	var avatarURL *string
	if row.AvatarUrl.Valid {
		avatarURL = &row.AvatarUrl.String
	}

	var dateOfBirth *time.Time
	if row.DateOfBirth.Valid {
		date := row.DateOfBirth.Time
		dateOfBirth = &date
	}

	return &account.Profile{
		UserID:      account.UserID{Value: row.UserID.Bytes},
		FullName:    row.FullName,
		Gender:      account.Gender(row.Gender),
		PhoneNumber: phoneNumber,
		AvatarURL:   avatarURL,
		DateOfBirth: dateOfBirth,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}

func toInfraUser(u *account.User) repository.InsertUserParams {
	return repository.InsertUserParams{
		ID:              common.ToPgUUID(u.ID.Value),
		Email:           u.Email,
		EmailVerifiedAt: common.ToPgTimeStampZ(u.EmailVerifiedAt),
		Status:          string(u.Status),
		Roles:           common.ToStringRoles(u.Roles),
		CreatedAt:       common.ToPgTimeStampZ(&u.CreatedAt),
		UpdatedAt:       common.ToPgTimeStampZ(&u.UpdatedAt),
		DeletedAt:       common.ToPgTimeStampZ(u.DeletedAt),
	}
}

func toInfraCredential(c *account.Credential) repository.InsertCredentialParams {
	return repository.InsertCredentialParams{
		UserID:            common.ToPgUUID(c.UserID.Value),
		PasswordHash:      c.PasswordHash,
		PasswordVersion:   int32(c.PasswordVersion),
		PasswordUpdatedAt: common.ToPgTimeStampZ(&c.PasswordUpdatedAt),
	}
}

func toInfraProfile(p *account.Profile) repository.InsertProfileParams {
	return repository.InsertProfileParams{
		UserID:      common.ToPgUUID(p.UserID.Value),
		FullName:    p.FullName,
		Gender:      string(p.Gender),
		PhoneNumber: common.ToPgText(p.PhoneNumber),
		AvatarUrl:   common.ToPgText(p.AvatarURL),
		DateOfBirth: common.ToPgDate(p.DateOfBirth),
		CreatedAt:   common.ToPgTimeStampZ(&p.CreatedAt),
		UpdatedAt:   common.ToPgTimeStampZ(&p.UpdatedAt),
	}
}

func toUpdateUserInfra(u *account.User) repository.UpdateUserParams {
	now := time.Now().UTC()
	return repository.UpdateUserParams{
		ID:              common.ToPgUUID(u.ID.Value),
		Email:           u.Email,
		Status:          string(u.Status),
		EmailVerifiedAt: common.ToPgTimeStampZ(u.EmailVerifiedAt),
		Roles:           common.ToStringRoles(u.Roles),
		UpdatedAt:       common.ToPgTimeStampZ(&now),
		DeletedAt:       common.ToPgTimeStampZ(u.DeletedAt),
	}
}
