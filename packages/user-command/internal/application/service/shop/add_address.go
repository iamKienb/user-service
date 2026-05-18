package shop

import (
	"context"
	"user-command-module/internal/application/command/add_shop_address"
	"user-command-module/internal/domain/auth"
	domain_shared "user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
)

func (s *shopService) AddAddress(ctx context.Context, cmd add_shop_address.Command) (*add_shop_address.Result, error) {
	userRoleIDs, err := s.getUserRoles(ctx, cmd.ShopID, cmd.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.authorizer.Authorize(auth.ActionShopManageAddress, userRoleIDs); err != nil {
		return nil, s.wrapError(err)
	}

	addressType, err := domain_shared.ValidateEnum[shop.AddressTypeEnum](cmd.Type, shop.ErrAddressTypeInvalid)
	if err != nil {
		return nil, s.wrapError(err)
	}

	address := shop.NewShopAddress(shop.AddressParams{
		UserID:      cmd.UserID,
		ShopID:      cmd.ShopID,
		CountryID:   cmd.Country.ID,
		CityID:      cmd.City.ID,
		DistrictID:  cmd.District.ID,
		WardID:      cmd.Ward.ID,
		Type:        addressType,
		ContactName: cmd.ContactName,
		PhoneNumber: cmd.PhoneNumber,
		AddressLine: cmd.AddressLine,
	})

	if err := s.shopRepo.SaveAddress(ctx, address); err != nil {
		return nil, s.wrapError(err)
	}

	return &add_shop_address.Result{
		ShopAddressID: address.ID.String(),
	}, nil
}
