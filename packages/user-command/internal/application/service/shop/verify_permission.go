package shop

import (
	"context"
	"user-command-module/internal/application/command/verify_permission"
)

func (s *shopService) VerifyPermission(ctx context.Context, cmd verify_permission.Command) (*verify_permission.Result, error) {
	userRoleIDs, err := s.GetRolesUser(ctx, cmd.ShopID, cmd.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.authorizer.Authorize(cmd.Action, userRoleIDs); err != nil {
		return &verify_permission.Result{
			IsAllowed:    false,
			ErrorMessage: err,
		}, nil
	}

	return &verify_permission.Result{
		IsAllowed:    true,
		ErrorMessage: nil,
	}, nil
}
