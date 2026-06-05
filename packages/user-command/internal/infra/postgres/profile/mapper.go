package profile

import (
	"time"
	"user-command-module/db/repository"
	domain_profile "user-command-module/internal/domain/profile"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toDomainProfile(row repository.UserProfile) *domain_profile.Profile {
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

	return &domain_profile.Profile{
		UserID:      shared.UserID(row.UserID.Bytes),
		FullName:    row.FullName,
		Gender:      domain_profile.GenderEnum(row.Gender),
		PhoneNumber: phoneNumber,
		AvatarURL:   avatarURL,
		DateOfBirth: dateOfBirth,
		CreatedAt:   row.CreatedAt.Time,
		UpdatedAt:   &row.UpdatedAt.Time,
	}
}

func toInfraProfile(p *domain_profile.Profile) repository.CreateUserProfileParams {
	return repository.CreateUserProfileParams{
		UserID:      conv.UUID(p.UserID),
		FullName:    p.FullName,
		Gender:      string(p.Gender),
		PhoneNumber: conv.Text(p.PhoneNumber),
		AvatarUrl:   conv.Text(p.AvatarURL),
		DateOfBirth: conv.Date(p.DateOfBirth),
		CreatedAt:   conv.TimeStampZ(&p.CreatedAt),
		UpdatedAt:   conv.TimeStampZ(p.UpdatedAt),
	}
}
