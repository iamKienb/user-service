package shop

import (
	"user-command-module/internal/application/service/shop/i18n"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shop"

	"github.com/iamKienb/go-core/app_error"
)

var shopErrorMap = app_error.ServiceErrorMap{
	shop.ErrShopSlugTaken:      {Kind: app_error.KindValidation, Msg: i18n.MsgSlugTaken},
	shop.ErrShopConflict:       {Kind: app_error.KindConflict, Msg: i18n.MsgCreateShopToFast},
	shop.ErrShopNotFound:       {Kind: app_error.KindNotFound, Msg: i18n.MsgShopNotFound},
	shop.ErrShopInvalid:        {Kind: app_error.KindValidation, Msg: i18n.MsgShopInvalid},
	shop.ErrAddressTypeInvalid: {Kind: app_error.KindValidation, Msg: i18n.MsgAddressTypeInvalid},

	auth.ErrActionNotDefined: {Kind: app_error.KindValidation, Msg: i18n.MsgActionInvalid},
	auth.ErrShopDenied:       {Kind: app_error.KindForbidden, Msg: i18n.MsgShopDenied},
	auth.ErrProductDenied:    {Kind: app_error.KindForbidden, Msg: i18n.MsgProductDenied},
	auth.ErrInventoryDenied:  {Kind: app_error.KindForbidden, Msg: i18n.MsgInventoryDenied},
}

func (s *shopService) wrapError(err error) error {
	return app_error.WrapError(err, shopErrorMap)
}
