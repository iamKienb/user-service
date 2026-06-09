package user

import (
	"context"
	get_user_address_by_id "user-command-module/internal/application/queries/get_address_by_id"
	domain_address "user-command-module/internal/domain/address"
	domain_auth "user-command-module/internal/domain/auth"
)

func (s *userService) GetAddress(ctx context.Context, qry get_user_address_by_id.Query) (*get_user_address_by_id.Result, error) {
	_, err := s.validateAndCheckActiveUser(ctx, qry.UserID)
	if err != nil {
		return nil, err
	}

	address, err := s.userAddressRepo.FindAddressByID(ctx, qry.UserAddressID)
	if err != nil {
		return nil, err
	}
	if address == nil {
		return nil, domain_address.ErrAddressNotFound
	}

	if address.UserID != qry.UserID {
		return nil, domain_auth.ErrAccessDenied
	}

	return &get_user_address_by_id.Result{
		UserAddressID: address.ID,
		UserID:        address.UserID,
		ReceiverName:  address.ReceiverName,
		PhoneNumber:   address.PhoneNumber,

		ProvinceID:   address.ProvinceID,
		ProvinceName: address.ProvinceName,

		WardID:   address.WardID,
		WardName: address.WardName,

		AddressLine: address.AddressLine,
		Label:       string(address.Label),
		IsDefault:   address.IsDefault,
	}, nil
}
