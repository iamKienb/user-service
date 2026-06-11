package otp

import (
	"user-command-module/internal/application/services/otp/i18n"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/user"

	"github.com/iamKienb/go-core/app_error"
)

var otpErrorMap = app_error.ServiceErrorMap{
	user.ErrUserNotFound: {Kind: app_error.KindNotFound, Msg: i18n.MsgUserNotFound},

	auth.ErrOTPInvalid:     {Kind: app_error.KindValidation, Msg: i18n.MsgOTPInvalid},
	auth.ErrSessionInvalid: {Kind: app_error.KindUnauthorized, Msg: i18n.MsgSessionInvalid},

	auth.ErrOTPExpired:  {Kind: app_error.KindUnauthorized, Msg: i18n.MsgOTPExpired},
	auth.ErrOTPMaxRetry: {Kind: app_error.KindValidation, Msg: i18n.MsgOTPMaxRetry},
	auth.ErrResendLimit: {Kind: app_error.KindValidation, Msg: i18n.MsgResendLimit},
}

func toApplicationError(err error) error {
	return app_error.WrapError(err, otpErrorMap)
}
