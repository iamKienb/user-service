package otp

import (
	"shopify-user-command-module/internal/application/service/otp/i18n"
	"shopify-user-command-module/internal/domain/identity"

	"github.com/iamKienb/shopify-go-platform/app_error"
)

var otpErrorMap = app_error.ServiceErrorMap{
	identity.ErrOTPInvalid:     {Kind: app_error.KindValidation, Msg: i18n.MsgOTPInvalid},
	identity.ErrSessionInvalid: {Kind: app_error.KindUnauthorized, Msg: i18n.MsgSessionInvalid},

	identity.ErrOTPExpired:  {Kind: app_error.KindUnauthorized, Msg: i18n.MsgOTPExpired},
	identity.ErrOTPMaxRetry: {Kind: app_error.KindValidation, Msg: i18n.MsgOTPMaxRetry},
	identity.ErrResendLimit: {Kind: app_error.KindValidation, Msg: i18n.MsgResendLimit},
}

func (s *otpService) wrapError(err error) error {
	return app_error.WrapError(err, otpErrorMap)
}
