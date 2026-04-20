package user

import (
	"shopify-user-command-module/internal/application/service/user/i18n"
	"shopify-user-command-module/internal/domain/identity"

	"github.com/iamKienb/shopify-go-platform/app_error"
)

var userErrorMap = app_error.ServiceErrorMap{
	identity.ErrEmailTaken:   {Kind: app_error.KindConflict, Msg: i18n.MsgEmailTaken},
	identity.ErrUserNotFound: {Kind: app_error.KindNotFound, Msg: i18n.MsgUserNotFound},
	identity.ErrUserInvalid:  {Kind: app_error.KindValidation, Msg: i18n.MsgUserInvalid},
}

func (s *userService) wrapError(err error) error {
	return app_error.WrapError(err, userErrorMap)
}
