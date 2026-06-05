package address

import (
	"user-command-module/db/repository"
	domain_address "user-command-module/internal/domain/address"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toInfraUserAddress(params *domain_address.UserAddress) repository.CreateUserAddressParams {
	return repository.CreateUserAddressParams{
		ID:     conv.UUID(params.ID),
		UserID: conv.UUID(params.UserID),

		CountryID:   params.CountryID,
		CountryName: params.CountryName,

		ProvinceID:   params.ProvinceID,
		ProvinceName: params.ProvinceName,

		WardID:   params.WardID,
		WardName: params.WardName,

		AddressLine:  params.AddressLine,
		ReceiverName: params.ReceiverName,
		PhoneNumber:  params.PhoneNumber,
		Label:        string(params.Label),
		IsDefault:    params.IsDefault,

		CreatedAt: conv.TimeStampZ(&params.CreatedAt),
		UpdatedAt: conv.TimeStampZ(params.UpdatedAt),
	}
}

func toDomainUserAddress(row repository.UserAddress) *domain_address.UserAddress {
	return &domain_address.UserAddress{
		ID:     shared.UserAddressID(row.ID.Bytes),
		UserID: shared.UserID(row.UserID.Bytes),

		CountryID:   row.CountryID,
		CountryName: row.CountryName,

		ProvinceID:   row.ProvinceID,
		ProvinceName: row.ProvinceName,

		WardID:   row.WardID,
		WardName: row.WardName,

		AddressLine:  row.AddressLine,
		ReceiverName: row.ReceiverName,
		PhoneNumber:  row.PhoneNumber,
		Label:        domain_address.LabelEnum(row.Label),
		IsDefault:    row.IsDefault,

		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: &row.UpdatedAt.Time,
	}
}
