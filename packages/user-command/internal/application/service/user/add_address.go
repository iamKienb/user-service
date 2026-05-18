package user

import (
	"context"
	"user-command-module/internal/application/command/add_user_address"
)

func (s *userService) AddAddress(ctx context.Context, cmd add_user_address.Command) (*add_user_address.Result, error) {
	// parseUserID, err := shared.ParseToRawID[domain_shared.UserID](cmd.UserID)
	// agg, err := s.accountRepo.LoadAggByID(ctx, parseUserID)
	return nil, nil
}
