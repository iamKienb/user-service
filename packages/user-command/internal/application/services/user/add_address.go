package user

import (
	"context"
	"user-command-module/internal/application/commands/add_user_address"
	"user-command-module/internal/domain/address"
	domain_shared "user-command-module/internal/domain/shared"
)

func (s *userService) AddAddress(ctx context.Context, cmd add_user_address.Command) (*add_user_address.Result, error) {
	_, err := s.validateAndCheckActiveUser(ctx, cmd.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	label := domain_shared.ValidateEnum[address.LabelEnum](cmd.Label)
	if label == nil {
		return nil, s.wrapError(address.ErrLabelInvalid)
	}

	newAddress := address.NewUserAddress(address.NewUserAddressParams{
		UserID:       cmd.UserID,
		CountryID:    cmd.Country.ID,
		CountryName:  cmd.Country.Name,
		ProvinceID:   cmd.Province.ID,
		ProvinceName: cmd.Province.Name,
		WardID:       cmd.Ward.ID,
		WardName:     cmd.Ward.Name,
		AddressLine:  cmd.AddressLine,
		ReceiverName: cmd.ReceiverName,
		PhoneNumber:  cmd.PhoneNumber,
		Label:        *label,
		IsDefault:    cmd.IsDefault,
	})

	if err := s.userAddressRepo.CreateUserAddress(ctx, newAddress); err != nil {
		return nil, s.wrapError(err)
	}

	return &add_user_address.Result{
		UserAddressID: newAddress.ID,
	}, nil
}
