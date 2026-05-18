package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

type ShopCreatedEvent struct {
	ShopID      shared.ShopID
	OwnerID     shared.UserID
	Slug        string
	Status      string
	Description *string
	LogoUrl     *string
	BannerUrl   *string
	CreatedBy   shared.UserID
	CreatedAt   time.Time
}

func (e ShopCreatedEvent) EventName() string {
	return "user-service.shop.created"
}

func (e ShopCreatedEvent) IntegrationPayload() map[string]interface{} {
	return map[string]interface{}{
		"shop_id":     e.ShopID.String(),
		"owner_id":    e.OwnerID.String(),
		"slug":        e.Slug,
		"status":      e.Status,
		"description": *e.Description,
		"logo_url":    *e.LogoUrl,
		"banner_url":  *e.BannerUrl,
		"created_by":  e.CreatedBy,
		"created_at":  e.CreatedAt,
	}
}

type RoleEvent struct {
	ID   shared.RoleID
	Code string
	Name string
}

type MemberAddedEvent struct {
	ShopID      shared.ShopID
	MemberID    shared.UserID
	MemberName  string
	AddedBy     shared.UserID
	NameAddedBy string
	Roles       []RoleEvent
	JoinedAt    time.Time
}

func (e MemberAddedEvent) EventName() string {
	return "shop.member_added"
}

func (e MemberAddedEvent) IntegrationPayload() map[string]interface{} {
	rolesMap := make([]map[string]interface{}, 0, len(e.Roles))
	for _, r := range e.Roles {
		rolesMap = append(rolesMap, map[string]interface{}{
			"id":   int(r.ID),
			"code": r.Code,
			"name": r.Name,
		})
	}

	return map[string]interface{}{
		"shop_id":       e.ShopID.String(),
		"member_id":     e.MemberID.String(),
		"member_name":   e.MemberName,
		"added_by":      e.AddedBy.String(),
		"name_added_by": e.NameAddedBy,
		"joined_at":     e.JoinedAt,
		"roles":         rolesMap,
	}
}
