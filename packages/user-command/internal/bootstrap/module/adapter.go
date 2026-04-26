package module

import (
	"log/slog"
	"net/http"

	otpadapter "shopify-user-command-module/internal/adapter/otp"
	useradapter "shopify-user-command-module/internal/adapter/user"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/iamKienb/shopify-go-api/gen/otp/otpconnect"
	"github.com/iamKienb/shopify-go-api/gen/user/userconnect"
	authx "github.com/iamKienb/shopify-go-platform/middleware/auth"
	observabilityx "github.com/iamKienb/shopify-go-platform/middleware/observability"
)

type AdapterModule struct {
	Mux *http.ServeMux
}

func NewAdapterModule(app *ApplicationModule, logger *slog.Logger) *AdapterModule {
	var interceptors []connect.Interceptor

	tracingInterceptor, err := observabilityx.TracingInterceptor()
	if err != nil {
		logger.Error("failed to initialize tracing interceptor", slog.Any("error", err))
	} else {
		interceptors = append(interceptors, tracingInterceptor)
	}

	interceptors = append(interceptors,
		observabilityx.RecoveryInterceptor(logger),
		authx.AuthInternalInterceptor(),
		observabilityx.LoggingInterceptor(logger),
		observabilityx.ValidationRequestInterceptor(),
		observabilityx.ErrorResponseInterceptor(),
	)

	allInterceptors := connect.WithInterceptors(interceptors...)

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		userconnect.UserCommandServiceName,
		otpconnect.OTPCommandServiceName,
	)

	userServer := useradapter.NewUserServer(app.RegisterExecutor, app.LoginExecutor)
	otpServer := otpadapter.NewOTPServer(app.VerifyExecutor, app.ResendExecutor)

	mux.Handle(userconnect.NewUserCommandServiceHandler(userServer, allInterceptors))
	mux.Handle(otpconnect.NewOTPCommandServiceHandler(otpServer, allInterceptors))

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return &AdapterModule{Mux: mux}
}
