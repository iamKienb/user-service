package user

import (
	"user-command-module/internal/application/services/user/i18n"
	"user-command-module/internal/domain/address"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/profile"
	domain_user "user-command-module/internal/domain/user"

	"github.com/iamKienb/go-core/app_error"
)

var userErrorMap = app_error.ServiceErrorMap{
	domain_user.ErrEmailEmpty:         {Kind: app_error.KindValidation, Msg: i18n.MsgEmailEmpty},
	domain_user.ErrEmailInvalid:       {Kind: app_error.KindValidation, Msg: i18n.MsgEmailInvalid},
	domain_user.ErrEmailTaken:         {Kind: app_error.KindConflict, Msg: i18n.MsgEmailTaken},
	domain_user.ErrUserNotFound:       {Kind: app_error.KindNotFound, Msg: i18n.MsgUserNotFound},
	domain_user.ErrUserInvalid:        {Kind: app_error.KindValidation, Msg: i18n.MsgUserInvalid},
	domain_user.ErrCredentialNotFound: {Kind: app_error.KindNotFound, Msg: i18n.MsgCredentialNotFound},
	domain_user.ErrUserNotActive:      {Kind: app_error.KindForbidden, Msg: i18n.MsgUserNotActive},

	profile.ErrNameEmpty:     {Kind: app_error.KindValidation, Msg: i18n.MsgNameEmpty},
	profile.ErrNameTooLong:   {Kind: app_error.KindValidation, Msg: i18n.MsgNameTooLong},
	profile.ErrGenderInvalid: {Kind: app_error.KindValidation, Msg: i18n.MsgGenderInvalid},

	address.ErrLabelInvalid:    {Kind: app_error.KindValidation, Msg: i18n.MsgLabelInvalid},
	address.ErrAddressNotFound: {Kind: app_error.KindNotFound, Msg: i18n.MsgAddressNotFound},

	auth.ErrInvalidCredentials: {Kind: app_error.KindUnauthorized, Msg: i18n.MsgInvalidCredentials},
	auth.ErrAccountLocked:      {Kind: app_error.KindForbidden, Msg: i18n.MsgAccountLocked},
	auth.ErrAccessDenied:       {Kind: app_error.KindUnauthorized, Msg: i18n.MsgAccessDenied},
}

func toMapError(err error) error {
	return app_error.WrapError(err, userErrorMap)
}
