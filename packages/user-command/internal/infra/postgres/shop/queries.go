package shop

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"

	"github.com/jackc/pgx/v5"
)

func (r *shopRepository) LoadAggByShopID(ctx context.Context, shopID shared.ShopID) (*shop.Aggregate, error) {
	return nil, nil
}

func (r *shopRepository) LoadAggByUserID(ctx context.Context, shopID shared.ShopID) (*shop.Aggregate, error) {
	return nil, nil
}

func (r *shopRepository) GetUserRolesInShop(ctx context.Context, shopID shared.ShopID, userID shared.UserID) (*shop.MemberPermission, error) {
	row, err := r.getQuerier(ctx).GetUserRolesInShop(ctx, toInfraGetUserRoleInShop(shopID, userID))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("infra: get user roles in shop: %w", err)
	}

	return toDomainMemberPermission(row), nil
}

func (r *shopRepository) CheckSlugExists(ctx context.Context, slug string) (bool, error) {
	_, err := r.queries.CountBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return true, nil
		}
		return false, fmt.Errorf("infra:postgres: count by slug: %w", err)
	}

	return false, nil
}
