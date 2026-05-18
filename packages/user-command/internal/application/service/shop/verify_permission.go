package shop

import (
	"context"
	"user-command-module/internal/application/command/verify_permission"
	"user-command-module/internal/application/service/shop/i18n"
	"user-command-module/internal/domain/auth"
)

func (s *shopService) VerifyPermission(ctx context.Context, cmd verify_permission.Command) (*verify_permission.Result, error) {
	userRoleIDs, err := s.getUserRoles(ctx, cmd.ShopID, cmd.UserID)
	if err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.authorizer.Authorize(cmd.Action, userRoleIDs); err != nil {
		return &verify_permission.Result{
			IsAllowed:    false,
			ErrorMessage: permissionMessage(err),
		}, nil
	}

	return &verify_permission.Result{
		IsAllowed:    true,
		ErrorMessage: "",
	}, nil
}

func permissionMessage(err error) string {
	switch err {
	case auth.ErrActionNotDefined:
		return i18n.MsgActionInvalid
	case auth.ErrShopDenied:
		return i18n.MsgShopDenied
	case auth.ErrProductDenied:
		return i18n.MsgProductDenied
	case auth.ErrInventoryDenied:
		return i18n.MsgInventoryDenied
	default:
		return ""
	}
}
