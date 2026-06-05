package module

import (
	"user-query-module/internal/application/queries/get_user_detail"
	"user-query-module/internal/application/queries/get_user_profile"
	"user-query-module/internal/application/queries/list_user_addresses"
	"user-query-module/internal/application/queries/search_users"
	"user-query-module/internal/application/service"
)

type ApplicationModule struct {
	GetUserDetailExecutor     get_user_detail.Executor
	GetUserProfileExecutor    get_user_profile.Executor
	ListUserAddressesExecutor list_user_addresses.Executor
	SearchUsersExecutor       search_users.Executor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	userQueryService := service.NewQueryService(infra.ESService)

	return &ApplicationModule{
		GetUserDetailExecutor:     get_user_detail.NewHandler(userQueryService),
		GetUserProfileExecutor:    get_user_profile.NewHandler(userQueryService),
		ListUserAddressesExecutor: list_user_addresses.NewHandler(userQueryService),
		SearchUsersExecutor:       search_users.NewHandler(userQueryService),
	}
}
