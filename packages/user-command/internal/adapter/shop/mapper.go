package shop

import (
	"fmt"
	"user-command-module/internal/application/command/add_shop_address"
	"user-command-module/internal/application/command/assign_member"
	"user-command-module/internal/application/command/create_shop"
	"user-command-module/internal/application/command/verify_permission"
	"user-command-module/internal/domain/account"
	"user-command-module/internal/domain/shared"

	"github.com/iamKienb/api-contract/gen/shop"
)

func ToCreateShopCommand(userID string, userName string, req *shop.CreateShopRequest) (create_shop.Command, error) {
	var profile *create_shop.Profile
	if req.Profile != nil {
		profile = &create_shop.Profile{
			Description: req.Profile.Description,
			LogoUrl:     req.Profile.LogoUrl,
			BannerUrl:   req.Profile.BannerUrl,
		}
	}
	userParse, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return create_shop.Command{}, fmt.Errorf("invalid userID: %w", err)
	}

	return create_shop.Command{
		User: create_shop.User{
			ID:   userParse,
			Name: userName,
		},
		Name:    req.GetName(),
		Slug:    req.GetSlug(),
		Profile: profile,
	}, nil
}

func ToCreateShopResponse(result *create_shop.Result) *shop.CreateShopResponse {
	return &shop.CreateShopResponse{
		ShopId: result.ShopID,
	}
}

func ToAssignMemberCommand(userID string, userName string, req *shop.AssignMemberRolesRequest) (assign_member.Command, error) {
	shopID, err := shared.ParseToRawID[shared.ShopID](req.GetShopId())
	if err != nil {
		return assign_member.Command{}, fmt.Errorf("invalid shopID: %w", err)
	}
	userParse, err := shared.ParseToRawID[shared.UserID](userID)
	if err != nil {
		return assign_member.Command{}, fmt.Errorf("invalid userID: %w", err)
	}

	reqMemberRoles := req.GetMemberRoles()
	memberRoles := make([]assign_member.MemberRole, 0, len(reqMemberRoles))

	for _, memberRole := range reqMemberRoles {
		memberID, err := shared.ParseToRawID[shared.UserID](memberRole.GetId())
		if err != nil {
			return assign_member.Command{}, fmt.Errorf("invalid memberID: %w", err)
		}

		reqRoleIDs := memberRole.GetRoleIDs()
		roleIDs := make([]shared.RoleID, 0, len(reqRoleIDs))

		for _, roleID := range reqRoleIDs {
			roleIDs = append(roleIDs, shared.RoleID(roleID.GetId()))
		}

		memberRoles = append(memberRoles, assign_member.MemberRole{
			ID:      memberID,
			RoleIDs: roleIDs,
		})
	}

	return assign_member.Command{
		User: assign_member.User{
			ID:   userParse,
			Name: userName,
		},
		ShopID:      shopID,
		MemberRoles: memberRoles,
		Action:      req.GetAction(),
	}, nil
}

func ToAssignMemberResponse(result *assign_member.Result) *shop.AssignMemberRolesResponse {
	return &shop.AssignMemberRolesResponse{
		Success: result.Success,
	}
}

func ToAddAddressCommand(req *shop.AddShopAddressRequest) (add_shop_address.Command, error) {
	userID, err := shared.ParseToRawID[shared.UserID](req.GetUserId())
	if err != nil {
		return add_shop_address.Command{}, account.ErrUserInvalid
	}
	shopID, err := shared.ParseToRawID[shared.ShopID](req.GetShopId())
	if err != nil {
		return add_shop_address.Command{}, fmt.Errorf("invalid shopID: %w", err)
	}

	return add_shop_address.Command{
		UserID: userID,
		ShopID: shopID,

		Country: add_shop_address.LocationInfo{
			ID:   int(req.Country.GetId()),
			Name: req.Country.GetName(),
		},
		City: add_shop_address.LocationInfo{
			ID:   int(req.City.GetId()),
			Name: req.City.GetName(),
		},
		District: add_shop_address.LocationInfo{
			ID:   int(req.District.GetId()),
			Name: req.District.GetName(),
		},
		Ward: add_shop_address.LocationInfo{
			ID:   int(req.Ward.GetId()),
			Name: req.Ward.GetName(),
		},

		AddressLine: req.GetAddressLine(),
		ContactName: req.GetContactName(),
		PhoneNumber: req.GetPhoneNumber(),
		Type:        req.GetType(),
	}, nil
}

func ToAddAddressResponse(result *add_shop_address.Result) *shop.AddShopAddressResponse {
	return &shop.AddShopAddressResponse{
		AddressId: result.ShopAddressID,
	}
}

func ToVerifyPermissionCommand(req *shop.VerifyPermissionRequest) (verify_permission.Command, error) {
	userID, err := shared.ParseToRawID[shared.UserID](req.GetUserId())
	if err != nil {
		return verify_permission.Command{}, account.ErrUserInvalid
	}

	shopID, err := shared.ParseToRawID[shared.ShopID](req.GetShopId())
	if err != nil {
		return verify_permission.Command{}, fmt.Errorf("invalid shopID: %w", err)
	}

	return verify_permission.Command{
		ShopID: shopID,
		UserID: userID,
		Action: req.GetAction(),
	}, nil

}

func ToVerifyPermissionResponse(result *verify_permission.Result) *shop.VerifyPermissionResponse {
	var errMsg string

	if result.ErrorMessage != nil {
		errMsg = result.ErrorMessage.Error()
	}

	return &shop.VerifyPermissionResponse{
		IsAllowed:    result.IsAllowed,
		ErrorMessage: errMsg,
	}
}
