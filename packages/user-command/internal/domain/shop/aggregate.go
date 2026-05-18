package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

const AggregateTypeShop = "SHOP"

type Aggregate struct {
	Shop    Shop
	Profile ShopProfile
	Events  []shared.DomainEvent
}

func NewAggregate(p AggregateParams) *Aggregate {
	now := time.Now().UTC()
	shopID := shared.NewID[shared.ShopID]()

	shop := Shop{
		ID:      shopID,
		OwnerID: p.UserID,
		Name:    p.Name,
		Slug:    p.Slug,
		Status:  StatusPending,

		CreatedBy: p.UserID,
		UpdatedBy: nil,

		CreatedAt: now,
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	profile := ShopProfile{
		ShopID:      shopID,
		Description: p.Description,
		LogoUrl:     p.LogoUrl,
		BannerUrl:   p.BannerUrl,

		CreatedBy: p.UserID,
		UpdatedBy: &p.UserID,

		CreatedAt: now,
		UpdatedAt: nil,
	}

	aggregate := &Aggregate{
		Shop:    shop,
		Profile: profile,
	}

	aggregate.Shop.EventEntity.AddEvent(ShopCreatedEvent{
		ShopID:      aggregate.Shop.ID,
		OwnerID:     p.UserID,
		Slug:        p.Slug,
		Status:      aggregate.Shop.Status,
		Description: p.Description,
		LogoUrl:     p.LogoUrl,
		BannerUrl:   p.BannerUrl,
		CreatedBy:   p.UserID,
		CreatedAt:   now,
	})

	return aggregate
}

func (a *Aggregate) FlushEvents() []shared.DomainEvent {
	var allEvents []shared.DomainEvent
	if len(a.Events) > 0 {
		allEvents = append(allEvents, a.Events...)
		a.Events = nil
	}

	allEvents = append(allEvents, a.Shop.EventEntity.Flush()...)
	return allEvents
}
