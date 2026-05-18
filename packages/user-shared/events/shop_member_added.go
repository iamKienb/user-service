package events

import "time"

const TopicShopMemberAdded = "user-service.user.registered"

type Role struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ShopMemberAdded struct {
	ShopID      string    `json:"shop_id"`
	MemberID    string    `json:"member_id"`
	MemberName  string    `json:"member_name"`
	JoinedAt    time.Time `json:"joined_at"`
	AddedBy     string    `json:"added_by"`
	NameAddedBy string    `json:"name_added_by"`
	Roles       []Role    `json:"role"`
}

func (u ShopMemberAdded) EventName() string {
	return TopicShopMemberAdded
}
