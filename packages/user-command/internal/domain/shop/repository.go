package shop

import (
	"context"
	"user-command-module/internal/domain/shared"
)

type QueryRepository interface {
	GetUserRolesInShop(ctx context.Context, shopID shared.ShopID, userID shared.UserID) (*MemberPermission, error)
	CheckSlugExists(ctx context.Context, slug string) (bool, error)
}

type CommandRepository interface {
	SaveAggregate(ctx context.Context, agg *Aggregate) error
	UpsertMemberAggregate(ctx context.Context, memberAgg []*MemberAggregate) error
	ClearShopMemberRolesBatch(ctx context.Context, memberAgg []*MemberAggregate) error
	SaveAddress(ctx context.Context, addr *ShopAddress) error
}

type Repository interface {
	QueryRepository
	CommandRepository
}
