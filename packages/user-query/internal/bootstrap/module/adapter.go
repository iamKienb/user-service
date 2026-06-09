package module

import (
	"log/slog"
	"net/http"

	userAdapter "user-query-module/internal/adapter/user"

	"connectrpc.com/grpcreflect"
	"github.com/iamKienb/api-contract/gen/user/userconnect"
	observabilityx "github.com/iamKienb/go-core/middleware/observability"
)

type AdapterModule struct {
	Mux *http.ServeMux
}

func NewAdapterModule(app *ApplicationModule, logger *slog.Logger) *AdapterModule {
	server := userAdapter.NewQueryServer(
		app.GetUserDetailExecutor,
		app.GetUserProfileExecutor,
		app.ListUserAddressesExecutor,
		app.SearchUsersExecutor,
	)
	allInterceptors := observabilityx.InternalServerOption(logger)

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(userconnect.UserQueryServiceName)
	mux.Handle(userconnect.NewUserQueryServiceHandler(server, allInterceptors))
	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return &AdapterModule{Mux: mux}
}
