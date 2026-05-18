package user

import (
	"context"
	"user-command-module/internal/application/command/add_user_address"
	"user-command-module/internal/domain/account"
	domain_shared "user-command-module/internal/domain/shared"
)

func (s *userService) AddAddress(ctx context.Context, cmd add_user_address.Command) (*add_user_address.Result, error) {
	agg, err := s.accountRepo.LoadAggByID(ctx, cmd.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}
	if agg == nil {
		return nil, s.wrapError(account.ErrUserNotFound)
	}

	label, err := domain_shared.ValidateEnum[account.LabelEnum](cmd.Label, account.ErrAddressLabelInvalid)
	if err != nil {
		return nil, s.wrapError(err)
	}

	address := account.NewUserAddress(account.AddAddressParams{
		UserID:       cmd.UserID,
		CountryID:    cmd.Country.ID,
		CountryName:  cmd.Country.Name,
		CityID:       cmd.City.ID,
		CityName:     cmd.City.Name,
		DistrictID:   cmd.District.ID,
		DistrictName: cmd.District.Name,
		WardID:       cmd.Ward.ID,
		WardName:     cmd.Ward.Name,
		AddressLine:  cmd.AddressLine,
		ReceiverName: cmd.ReceiverName,
		PhoneNumber:  cmd.PhoneNumber,
		Label:        label,
		IsDefault:    cmd.IsDefault,
	})

	if err := s.accountRepo.SaveAddress(ctx, address); err != nil {
		return nil, s.wrapError(err)
	}

	return &add_user_address.Result{
		UserAddressID: address.ID.String(),
	}, nil
}
