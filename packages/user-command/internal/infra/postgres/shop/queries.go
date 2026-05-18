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
	_, err := r.getQuerier(ctx).CountBySlug(ctx, slug)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("infra:postgres: count by slug: %w", err)
	}

	return true, nil
}
