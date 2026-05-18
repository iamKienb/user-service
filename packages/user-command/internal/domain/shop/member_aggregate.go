package shop

import (
	"time"
	"user-command-module/internal/domain/shared"
)

const (
	RoleOwnerID     shared.RoleID = 1
	RoleManagerID   shared.RoleID = 2
	RoleCashierID   shared.RoleID = 3
	RoleWarehouseID shared.RoleID = 4
	RoleMarketingID shared.RoleID = 5
)

var RoleMetadata = map[shared.RoleID]struct {
	Code string
	Name string
}{
	RoleOwnerID:     {Code: "OWNER", Name: "Store Owner"},
	RoleManagerID:   {Code: "MANAGER", Name: "Store Manager"},
	RoleCashierID:   {Code: "CASHIER", Name: "Cashier"},
	RoleWarehouseID: {Code: "WAREHOUSE", Name: "Warehouse Staff"},
	RoleMarketingID: {Code: "MARKETING", Name: "Marketing Staff"},
}

type MemberAggregate struct {
	Member      ShopMember
	MemberRoles []ShopMemberRole
	Events      []shared.DomainEvent
}

func NewMemberAggregate(p MemberAggregateParams) *MemberAggregate {
	now := time.Now().UTC()

	member := ShopMember{
		ShopID:   p.ShopID,
		MemberID: p.MemberID,
		JoinedAt: now,
		AddedBy:  p.AddedBy,
	}

	memberRoles := make([]ShopMemberRole, 0, len(p.RoleIDs))
	eventRoles := make([]RoleEvent, 0, len(p.RoleIDs))

	for _, roleID := range p.RoleIDs {
		memberRoles = append(memberRoles, ShopMemberRole{
			ShopID:    p.ShopID,
			MemberID:  p.MemberID,
			RoleID:    roleID,
			UpdatedBy: p.AddedBy,
		})

		if meta, exists := RoleMetadata[roleID]; exists {
			eventRoles = append(eventRoles, RoleEvent{
				ID:   roleID,
				Code: meta.Code,
				Name: meta.Name,
			})
		}
	}

	aggregate := &MemberAggregate{
		Member:      member,
		MemberRoles: memberRoles,
	}

	aggregate.Member.EventEntity.AddEvent(MemberAddedEvent{
		ShopID:      p.ShopID,
		MemberID:    p.MemberID,
		MemberName:  p.MemberName,
		AddedBy:     p.AddedBy,
		NameAddedBy: p.NameAddedBy,
		JoinedAt:    now,
		Roles:       eventRoles,
	})

	return aggregate
}

func (a *MemberAggregate) FlushEvents() []shared.DomainEvent {
	var allEvents []shared.DomainEvent
	if len(a.Events) > 0 {
		allEvents = append(allEvents, a.Events...)
		a.Events = nil
	}

	allEvents = append(allEvents, a.Member.EventEntity.Flush()...)

	return allEvents
}
