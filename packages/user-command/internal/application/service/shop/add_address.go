package shop

import (
	"context"
	"user-command-module/internal/application/command/add_shop_address"
)

func (s *shopService) AddAddress(ctx context.Context, cmd add_shop_address.Command) (*add_shop_address.Result, error) {
	// parseUserID, err := shared.ParseToRawID[domain_shared.UserID](cmd.UserID)
	// agg, err := s.accountRepo.LoadAggByID(ctx, parseUserID)
	return nil, nil
}
