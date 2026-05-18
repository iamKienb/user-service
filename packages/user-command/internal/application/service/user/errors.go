package user

import (
	"user-command-module/internal/application/service/user/i18n"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/auth"

	"github.com/iamKienb/go-core/app_error"
)

var userErrorMap = app_error.ServiceErrorMap{
	account.ErrEmailEmpty:          {Kind: app_error.KindValidation, Msg: i18n.MsgEmailEmpty},
	account.ErrEmailInvalid:        {Kind: app_error.KindValidation, Msg: i18n.MsgEmailInvalid},
	account.ErrEmailTaken:          {Kind: app_error.KindConflict, Msg: i18n.MsgEmailTaken},
	account.ErrNameEmpty:           {Kind: app_error.KindValidation, Msg: i18n.MsgNameEmpty},
	account.ErrNameTooLong:         {Kind: app_error.KindValidation, Msg: i18n.MsgNameTooLong},
	account.ErrGenderInvalid:       {Kind: app_error.KindValidation, Msg: i18n.MsgGenderInvalid},
	account.ErrUserNotFound:        {Kind: app_error.KindNotFound, Msg: i18n.MsgUserNotFound},
	account.ErrUserInvalid:         {Kind: app_error.KindValidation, Msg: i18n.MsgUserInvalid},
	account.ErrCredentialNotFound:  {Kind: app_error.KindNotFound, Msg: i18n.MsgCredentialNotFound},
	account.ErrAddressLabelInvalid: {Kind: app_error.KindValidation, Msg: i18n.MsgAddressLabelInvalid},
	auth.ErrInvalidCredentials:     {Kind: app_error.KindUnauthorized, Msg: i18n.MsgInvalidCredentials},
	auth.ErrAccountLocked:          {Kind: app_error.KindForbidden, Msg: i18n.MsgAccountLocked},
	account.ErrUserNotActive:       {Kind: app_error.KindForbidden, Msg: i18n.MsgUserNotActive},
}

func (s *userService) wrapError(err error) error {
	return app_error.WrapError(err, userErrorMap)
}
