package address

import (
	"user-command-module/db/repository"
	domain_address "user-command-module/internal/domain/address"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/go-core/postgres/conv"
)

func toInfraUserAddress(params *domain_address.UserAddress) repository.SaveUserAddressParams {
	return repository.SaveUserAddressParams{
		ID:     conv.UUID(params.ID),
		UserID: conv.UUID(params.UserID),

		CountryID:  int32(params.CountryID),
		CityID:     int32(params.CityID),
		DistrictID: int32(params.DistrictID),
		WardID:     int32(params.WardID),

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

		CountryID:  int(row.CountryID),
		CityID:     int(row.CityID),
		DistrictID: int(row.DistrictID),
		WardID:     int(row.WardID),

		AddressLine:  row.AddressLine,
		ReceiverName: row.ReceiverName,
		PhoneNumber:  row.PhoneNumber,
		Label:        domain_address.LabelEnum(row.Label),
		IsDefault:    row.IsDefault,

		CreatedAt: row.CreatedAt.Time,
		UpdatedAt: &row.UpdatedAt.Time,
	}
}
