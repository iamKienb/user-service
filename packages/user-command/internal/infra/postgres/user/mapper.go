package user

import (
	"shopify-user-command-module/db/repository"
	"shopify-user-command-module/internal/domain/identity"
	"shopify-user-command-module/internal/infra/common"
	"time"
)

// ToDomain
func toDomainUser(row repository.User) *identity.User {
	return &identity.User{
		ID:        identity.UserID{Value: row.ID.Bytes},
		Email:     row.Email,
		Status:    identity.UserStatus(row.Status),
		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}

func toDomainCredential(row repository.Credential) *identity.Credential {
	return &identity.Credential{
		UserID:            identity.UserID{Value: row.UserID.Bytes},
		PasswordVersion:   int(row.PasswordVersion),
		PasswordUpdatedAt: row.PasswordUpdatedAt.Time,
	}
}

func toDomainProfile(row repository.Profile) *identity.Profile {
	return &identity.Profile{
		UserID:      identity.UserID{Value: row.UserID.Bytes},
		FullName:    row.FullName,
		Gender:      row.Gender,
		PhoneNumber: &row.PhoneNumber.String,
		AvatarURL:   &row.AvatarUrl.String,
		DateOfBirth: &row.DateOfBirth.Time,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   row.UpdatedAt.Time,
	}
}

// ToInfra
func toInfraUser(u *identity.User) repository.InsertUserParams {
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

func toInfraCredential(c *identity.Credential) repository.InsertCredentialParams {
	return repository.InsertCredentialParams{
		UserID:            common.ToPgUUID(c.UserID.Value),
		PasswordHash:      c.PasswordHash,
		PasswordVersion:   int32(c.PasswordVersion),
		PasswordUpdatedAt: common.ToPgTimeStampZ(&c.PasswordUpdatedAt),
	}
}

func toInfraProfile(p *identity.Profile) repository.InsertProfileParams {
	return repository.InsertProfileParams{
		UserID:      common.ToPgUUID(p.UserID.Value),
		FullName:    p.FullName,
		Gender:      p.Gender,
		PhoneNumber: common.ToPgText(p.PhoneNumber),
		AvatarUrl:   common.ToPgText(p.AvatarURL),
		DateOfBirth: common.ToPgDate(p.DateOfBirth),
		CreatedAt:   common.ToPgTimeStampZ(&p.CreatedAt),
		UpdatedAt:   common.ToPgTimeStampZ(&p.UpdatedAt),
	}
}

// ToUpdateInfra
func toUpdateUserInfra(u *identity.User) repository.UpdateUserParams {
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

// func toUpdateCredentialInfra(c *identity.Credential) repository.UpdateCredentialParams {
// 	now := time.Now().UTC()
// 	return repository.UpdateCredentialParams{
// 		UserID:            common.ToPgUUID(c.UserID.Value),
// 		PasswordHash:      c.PasswordHash,
// 		PasswordVersion:   int32(c.PasswordVersion),
// 		PasswordUpdatedAt: common.ToPgTimeStampZ(&now),
// 	}
// }

// func toUpdateProfileInfra(p *identity.Profile) repository.UpdateProfileParams {
// 	now := time.Now().UTC()
// 	return repository.UpdateProfileParams{
// 		UserID:      common.ToPgUUID(p.UserID.Value),
// 		FullName:    p.FullName,
// 		Gender:      p.Gender,
// 		PhoneNumber: common.ToPgText(p.PhoneNumber),
// 		AvatarUrl:   common.ToPgText(p.AvatarURL),
// 		DateOfBirth: common.ToPgDate(p.DateOfBirth),
// 		UpdatedAt:   common.ToPgTimeStampZ(&now),
// 	}
// }
