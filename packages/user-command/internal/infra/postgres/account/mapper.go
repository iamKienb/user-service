package account

import (
	"time"
	"user-command-module/db/repository"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/shared"
	"user-shared-module/common"
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
		ID:              shared.UserID(row.ID.Bytes),
		Email:           row.Email,
		Status:          account.StatusEnum(row.Status),
		EmailVerifiedAt: emailVerifiedAt,
		Roles:           common.ToEnumSlice[account.RoleEnum](row.Roles),
		CreatedAt:       row.CreatedAt.Time,
		UpdatedAt:       &row.UpdatedAt.Time,
		DeletedAt:       deletedAt,
	}
}

func toDomainCredential(row repository.UserCredential) *account.UserCredential {
	return &account.UserCredential{
		UserID:            shared.UserID(row.UserID.Bytes),
		PasswordHash:      row.PasswordHash,
		PasswordVersion:   int(row.PasswordVersion),
		PasswordUpdatedAt: row.PasswordUpdatedAt.Time,
	}
}

func toDomainProfile(row repository.UserProfile) *account.UserProfile {
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

	return &account.UserProfile{
		UserID:      shared.UserID(row.UserID.Bytes),
		FullName:    row.FullName,
		Gender:      account.GenderEnum(row.Gender),
		PhoneNumber: phoneNumber,
		AvatarURL:   avatarURL,
		DateOfBirth: dateOfBirth,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   &row.UpdatedAt.Time,
	}
}

func toInfraUser(u *account.User) repository.SaveUserParams {
	return repository.SaveUserParams{
		ID:              common.ToPgUUID(u.ID),
		Email:           u.Email,
		EmailVerifiedAt: common.ToPgTimeStampZ(u.EmailVerifiedAt),
		Status:          string(u.Status),
		Roles:           common.ToStringSlice(u.Roles),
		CreatedAt:       common.ToPgTimeStampZ(&u.CreatedAt),
		UpdatedAt:       common.ToPgTimeStampZ(u.UpdatedAt),
		DeletedAt:       common.ToPgTimeStampZ(u.DeletedAt),
	}
}

func toInfraCredential(c *account.UserCredential) repository.SaveUserCredentialParams {
	return repository.SaveUserCredentialParams{
		UserID:            common.ToPgUUID(c.UserID),
		PasswordHash:      c.PasswordHash,
		PasswordVersion:   int32(c.PasswordVersion),
		PasswordUpdatedAt: common.ToPgTimeStampZ(&c.PasswordUpdatedAt),
	}
}

func toInfraProfile(p *account.UserProfile) repository.SaveUserProfileParams {
	return repository.SaveUserProfileParams{
		UserID:      common.ToPgUUID(p.UserID),
		FullName:    p.FullName,
		Gender:      string(p.Gender),
		PhoneNumber: common.ToPgText(p.PhoneNumber),
		AvatarUrl:   common.ToPgText(p.AvatarURL),
		DateOfBirth: common.ToPgDate(p.DateOfBirth),
		CreatedAt:   common.ToPgTimeStampZ(&p.CreatedAt),
		UpdatedAt:   common.ToPgTimeStampZ(p.UpdatedAt),
	}
}

func toUpdateUserInfra(u *account.User) repository.UpdateUserParams {
	now := time.Now().UTC()
	return repository.UpdateUserParams{
		ID:              common.ToPgUUID(u.ID),
		Email:           u.Email,
		Status:          string(u.Status),
		EmailVerifiedAt: common.ToPgTimeStampZ(u.EmailVerifiedAt),
		Roles:           common.ToStringSlice(u.Roles),
		UpdatedAt:       common.ToPgTimeStampZ(&now),
		DeletedAt:       common.ToPgTimeStampZ(u.DeletedAt),
	}
}

func toInfraUserAddress(p *account.UserAddress) repository.SaveUserAddressParams {
	return repository.SaveUserAddressParams{
		ID:     common.ToPgUUID(p.ID),
		UserID: common.ToPgUUID(p.UserID),

		CountryID:  int32(p.CountryID),
		CityID:     int32(p.CityID),
		DistrictID: int32(p.DistrictID),
		WardID:     int32(p.WardID),

		AddressLine:  p.AddressLine,
		ReceiverName: p.ReceiverName,
		PhoneNumber:  p.PhoneNumber,
		Label:        string(p.Label),
		IsDefault:    p.IsDefault,

		CreatedAt: common.ToPgTimeStampZ(&p.CreatedAt),
		UpdatedAt: common.ToPgTimeStampZ(&p.UpdatedAt),
	}
}

func toDomainAddress(row repository.UserAddress) *account.UserAddress {
	return &account.UserAddress{
		ID:     shared.UserAddressID(row.ID.Bytes),
		UserID: shared.UserID(row.UserID.Bytes),

		CountryID:  int(row.CountryID),
		CityID:     int(row.CityID),
		DistrictID: int(row.DistrictID),
		WardID:     int(row.WardID),

		AddressLine:  row.AddressLine,
		ReceiverName: row.ReceiverName,
		PhoneNumber:  row.PhoneNumber,
		Label:        account.LabelEnum(row.Label),
		IsDefault:    row.IsDefault,

		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: row.UpdatedAt.Time,
	}
}
