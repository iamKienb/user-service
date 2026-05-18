package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopProfile struct {
	ShopID      shared.ShopID
	Description *string
	LogoUrl     *string
	BannerUrl   *string

	CreatedBy shared.UserID
	UpdatedBy *shared.UserID

	CreatedAt time.Time
	UpdatedAt *time.Time
}
