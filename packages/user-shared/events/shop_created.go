package events

import "time"

const TopicShopCreated = "user-service.shop.created"

type ShopCreated struct {
	ShopID  string `json:"shop_id"`
	OwnerID string `json:"owner_id"`
	Slug    string `json:"slug"`
	Status  string `json:"status"`

	Description *string `json:"description"`
	LogoUrl     *string `json:"logo_url"`
	BannerUrl   *string `json:"banner_url"`

	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}

func (u ShopCreated) EventName() string {
	return TopicShopCreated
}
