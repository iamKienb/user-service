package otp

import (
	"shopify-user-command-module/internal/application/service/otp/i18n"
	"shopify-user-command-module/internal/domain/account"
	"shopify-user-command-module/internal/domain/auth"

	"github.com/iamKienb/shopify-go-platform/app_error"
)

var otpErrorMap = app_error.ServiceErrorMap{
	account.ErrUserNotFound: {Kind: app_error.KindNotFound, Msg: "user was not found"},

	auth.ErrOTPInvalid:     {Kind: app_error.KindValidation, Msg: i18n.MsgOTPInvalid},
	auth.ErrSessionInvalid: {Kind: app_error.KindUnauthorized, Msg: i18n.MsgSessionInvalid},

	auth.ErrOTPExpired:  {Kind: app_error.KindUnauthorized, Msg: i18n.MsgOTPExpired},
	auth.ErrOTPMaxRetry: {Kind: app_error.KindValidation, Msg: i18n.MsgOTPMaxRetry},
	auth.ErrResendLimit: {Kind: app_error.KindValidation, Msg: i18n.MsgResendLimit},
}

func (s *otpService) wrapError(err error) error {
	return app_error.WrapError(err, otpErrorMap)
}
