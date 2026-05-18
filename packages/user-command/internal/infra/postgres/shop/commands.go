package shop

import (
	"context"
	"fmt"
	"user-command-module/internal/domain/shop"
)

func (r *shopRepository) SaveAggregate(ctx context.Context, agg *shop.Aggregate) error {
	q := r.getQuerier(ctx)

	if err := q.SaveShop(ctx, toInfraShop(&agg.Shop)); err != nil {
		if r.IsDuplicateSlug(err) {
			return shop.ErrShopSlugTaken
		}
		return fmt.Errorf("infra: save shop failed: %w", err)
	}

	if err := q.SaveShopProfile(ctx, toInfraProfile(&agg.Profile)); err != nil {
		return fmt.Errorf("infra: save profile failed: %w", err)
	}

	return nil
}

func (r *shopRepository) UpsertMemberAggregate(ctx context.Context, members []*shop.MemberAggregate) error {
	if len(members) == 0 {
		return nil
	}

	q := r.getQuerier(ctx)
	memberModels, roleModels := toInfraMemberRoleBatch(members)

	if err := q.AddShopMembersBatch(ctx, memberModels); err != nil {
		return fmt.Errorf("infra: add shop member batch failed: %w", err)
	}

	if err := q.AssignShopMemberRolesBatch(ctx, roleModels); err != nil {
		return fmt.Errorf("infra: assign shop member roles batch failed: %w", err)
	}

	return nil
}

func (r *shopRepository) ClearShopMemberRolesBatch(ctx context.Context, members []*shop.MemberAggregate) error {
	q := r.getQuerier(ctx)

	if err := q.ClearShopMemberRolesBatch(ctx, toInfraClearMemberRoles(members)); err != nil {
		return fmt.Errorf("infra: clear shop member batch failed: %w", err)
	}

	return nil
}

func (r *shopRepository) SaveAddress(ctx, addr *shop.ShopAddress) error {
	return nil
}
