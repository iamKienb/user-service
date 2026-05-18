package auth

import (
	"user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
)

var (
	roleOwnerManagerMarketing = []shared.RoleID{shop.RoleOwnerID, shop.RoleManagerID, shop.RoleMarketingID}
	roleOwnerManager          = []shared.RoleID{shop.RoleOwnerID, shop.RoleManagerID}
)

type IAuthorizer interface {
	Authorize(action string, roleIDs []shared.RoleID) error
}

type PermissionRule struct {
	AllowedRoles []shared.RoleID
	Error        error
}

type Authorizer struct {
	rules map[string]PermissionRule
	Error error
}

func NewAuthorizer() IAuthorizer {
	return &Authorizer{
		rules: map[string]PermissionRule{
			"product:create": {
				AllowedRoles: roleOwnerManagerMarketing,
				Error:        ErrProductDenied,
			},
			"shop:add_member": {
				AllowedRoles: roleOwnerManager,
				Error:        ErrShopDenied,
			},
		},
	}
}

func (a *Authorizer) Authorize(action string, roleIDs []shared.RoleID) error {
	rule, exists := a.rules[action]
	if !exists {
		return ErrorAction
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
