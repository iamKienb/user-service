package shop

import (
	"user-command-module/db/repository"
	"user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
	"user-shared-module/common"

	"github.com/jackc/pgx/v5/pgtype"
)

func toInfraClearMemberRoles(members []*shop.MemberAggregate) repository.ClearShopMemberRolesBatchParams {
	if len(members) == 0 {
		return repository.ClearShopMemberRolesBatchParams{}
	}
	memberIDs := make([]pgtype.UUID, 0, len(members))

	for _, agg := range members {
		if agg == nil {
			continue
		}

		memberID := common.ToPgUUID(agg.Member.MemberID)
		memberIDs = append(memberIDs, memberID)
	}

	return repository.ClearShopMemberRolesBatchParams{
		ShopID:    common.ToPgUUID(members[0].Member.ShopID),
		MemberIds: memberIDs,
	}
}

func toDomainMemberPermission(rows []repository.GetUserRolesInShopRow) *shop.MemberPermission {
	if len(rows) == 0 {
		return nil
	}

	roleIds := make([]shared.RoleID, 0, len(rows))

	for _, row := range rows {
		if !row.RoleID.Valid {
			continue
		}
		roleIds = append(roleIds, shared.RoleID(row.RoleID.Int32))
	}

	return &shop.MemberPermission{
		ShopID:  rows[0].ShopID.Bytes,
		RoleIDs: roleIds,
	}
}

func toNullableUUID(id *shared.UserID) pgtype.UUID {
	if id == nil {
		return pgtype.UUID{Valid: false}
	}

	return common.ToPgUUID(*id)
}

func toInfraGetUserRoleInShop(shopID shared.ShopID, userID shared.UserID) repository.GetUserRolesInShopParams {
	return repository.GetUserRolesInShopParams{
		ID:       common.ToPgUUID(shopID),
		MemberID: common.ToPgUUID(userID),
	}
}

func toInfraShop(shop *shop.Shop) repository.SaveShopParams {
	return repository.SaveShopParams{
		ID:        common.ToPgUUID(shop.ID),
		OwnerID:   common.ToPgUUID(shop.OwnerID),
		Name:      shop.Name,
		Slug:      shop.Slug,
		Status:    string(shop.Status),
		CreatedBy: common.ToPgUUID(shop.CreatedBy),
		UpdatedBy: toNullableUUID(shop.UpdatedBy),
		CreatedAt: common.ToPgTimeStampZ(&shop.CreatedAt),
		UpdatedAt: common.ToPgTimeStampZ(shop.UpdatedAt),
	}
}

func toInfraProfile(profile *shop.ShopProfile) repository.SaveShopProfileParams {
	return repository.SaveShopProfileParams{
		ShopID:      common.ToPgUUID(profile.ShopID),
		Description: common.ToPgText(profile.Description),
		LogoUrl:     common.ToPgText(profile.LogoUrl),
		BannerUrl:   common.ToPgText(profile.BannerUrl),
		CreatedBy:   common.ToPgUUID(profile.CreatedBy),
		UpdatedBy:   toNullableUUID(profile.UpdatedBy),
		CreatedAt:   common.ToPgTimeStampZ(&profile.CreatedAt),
		UpdatedAt:   common.ToPgTimeStampZ(profile.UpdatedAt),
	}
}

func toInfraMemberRoleBatch(members []*shop.MemberAggregate) (repository.AddShopMembersBatchParams, repository.AssignShopMemberRolesBatchParams) {
	if len(members) == 0 {
		return repository.AddShopMembersBatchParams{}, repository.AssignShopMemberRolesBatchParams{}
	}

	firstMember := members[0].Member
	shopID := common.ToPgUUID(firstMember.ShopID)
	addedBy := common.ToPgUUID(firstMember.AddedBy)
	joinedAt := common.ToPgTimeStampZ(&firstMember.JoinedAt)

	var memberIDs []pgtype.UUID
	var roleMemberIDs []pgtype.UUID
	var roleIDs []int32

	for _, agg := range members {
		if agg == nil {
			continue
		}

		memberID := common.ToPgUUID(agg.Member.MemberID)
		memberIDs = append(memberIDs, memberID)

		for _, role := range agg.MemberRoles {
			roleMemberIDs = append(roleMemberIDs, memberID)
			roleIDs = append(roleIDs, int32(role.RoleID))
		}
	}

	membersParams := repository.AddShopMembersBatchParams{
		ShopID:    shopID,
		AddedBy:   addedBy,
		MemberIds: memberIDs,
		JoinedAt:  joinedAt,
	}

	memberRolesParams := repository.AssignShopMemberRolesBatchParams{
		ShopID:    shopID,
		MemberIds: roleMemberIDs,
		RoleIds:   roleIDs,
		UpdatedBy: addedBy,
	}

	return membersParams, memberRolesParams
}

func toInfraShopAddress(address *shop.ShopAddress) repository.SaveShopAddressParams {
	return repository.SaveShopAddressParams{
		ID:          common.ToPgUUID(address.ID),
		ShopID:      common.ToPgUUID(address.ShopID),
		CountryID:   int32(address.CountryID),
		CityID:      int32(address.CityID),
		DistrictID:  int32(address.DistrictID),
		WardID:      int32(address.WardID),
		AddressLine: address.AddressLine,
		ContactName: address.ContactName,
		PhoneNumber: address.PhoneNumber,
		Type:        string(address.Type),
		CreatedAt:   common.ToPgTimeStampZ(&address.CreatedAt),
		UpdatedAt:   common.ToPgTimeStampZ(address.UpdatedAt),
		CreatedBy:   toNullableUUID(address.CreatedBy),
		UpdatedBy:   toNullableUUID(address.UpdatedBy),
	}
}
