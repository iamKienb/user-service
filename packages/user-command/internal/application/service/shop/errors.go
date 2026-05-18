package shop

import (
	"user-command-module/internal/application/service/shop/i18n"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shop"

	"github.com/iamKienb/shopify-go-platform/app_error"
)

var shopErrorMap = app_error.ServiceErrorMap{
	shop.ErrShopSlugTaken: {Kind: app_error.KindValidation, Msg: i18n.MsgSlugTaken},
	shop.ErrShopConflict:  {Kind: app_error.KindConflict, Msg: i18n.MsgCreateShopToFast},
	shop.ErrShopNotFound:  {Kind: app_error.KindNotFound, Msg: i18n.MsgShopNotFound},

	shop.ErrShopSlugTaken: {Kind: app_error.KindValidation, Msg: i18n.MsgSlugTaken},
	shop.ErrShopConflict:  {Kind: app_error.KindConflict, Msg: i18n.MsgCreateShopToFast},

	auth.ErrShopDenied:    {Kind: app_error.KindForbidden, Msg: i18n.MsgShopDenied},
	auth.ErrProductDenied: {Kind: app_error.KindForbidden, Msg: i18n.MsgProductDenied},
}

func (s *shopService) wrapError(err error) error {
	return app_error.WrapError(err, shopErrorMap)
}
