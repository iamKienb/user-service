package auth

import (
	"user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
)

var (
	roleOwnerManagerMarketing = []shared.RoleID{shop.RoleOwnerID, shop.RoleManagerID, shop.RoleMarketingID}
	roleOwnerManager          = []shared.RoleID{shop.RoleOwnerID, shop.RoleManagerID}
)

const (
	ActionProductCreate     = "product:create"
	ActionShopAddMember     = "shop:add_member"
	ActionShopManageAddress = "shop:manage_addresses"
)

type Authorizer interface {
	Authorize(action string, roleIDs []shared.RoleID) error
}

type PermissionRule struct {
	AllowedRoles []shared.RoleID
	Error        error
}

type ruleAuthorizer struct {
	rules map[string]PermissionRule
}

func NewAuthorizer() Authorizer {
	return &ruleAuthorizer{
		rules: map[string]PermissionRule{
			ActionProductCreate: {
				AllowedRoles: roleOwnerManagerMarketing,
				Error:        ErrProductDenied,
			},
			ActionShopAddMember: {
				AllowedRoles: roleOwnerManager,
				Error:        ErrShopDenied,
			},
			ActionShopManageAddress: {
				AllowedRoles: roleOwnerManager,
				Error:        ErrShopDenied,
			},
		},
	}
}

func (a *ruleAuthorizer) Authorize(action string, roleIDs []shared.RoleID) error {
	rule, exists := a.rules[action]
	if !exists {
		return ErrActionNotDefined
	}

	userRoleMap := make(map[shared.RoleID]struct{}, len(roleIDs))
	for _, role := range roleIDs {
		userRoleMap[role] = struct{}{}
	}

	for _, allowedRole := range rule.AllowedRoles {
		if _, hasRole := userRoleMap[allowedRole]; hasRole {
			return nil
		}
	}

	return rule.Error
}
