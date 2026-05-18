package shop

import "errors"

var (
	ErrMaxAddressReached    = errors.New("only had create 10 address")
	ErrShopIsDisable        = errors.New("shop had disabled")
	ErrPickupAddressMissing = errors.New("pickup address is required")
	ErrReturnAddressMissing = errors.New("return address is required")

	ErrGenderInvalid = errors.New("gender_invalid")
	ErrLogoMissing   = errors.New("shop_logo_missing")
	ErrBannerMissing = errors.New("shop_banner_missing")

	ErrShopSlugTaken = errors.New("shop_slug_already_taken")
	ErrShopConflict  = errors.New("crete_shop_to_fast")

	ErrShopNotFound = errors.New("shop_not_found")
)
