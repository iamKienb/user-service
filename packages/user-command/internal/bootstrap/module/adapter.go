package module

import (
	"log/slog"
	"net/http"
	"shopify-user-command-module/contract/protogen/otp/otpconnect"
	"shopify-user-command-module/contract/protogen/user/userconnect"
	"shopify-user-command-module/internal/adapter/otp"
	"shopify-user-command-module/internal/adapter/user"

	"github.com/iamKienb/shopify-go-platform/middleware/observability"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
)

type AdapterModule struct {
	Mux *http.ServeMux
}

func NewAdapterModule(app *ApplicationModule, logger *slog.Logger) *AdapterModule {
	tracingInterceptor, err := observability.NewTracingInterceptor()
	if err != nil {
		logger.Error("failed to initialize tracing interceptor", slog.Any("error", err))
	}
	interceptors := connect.WithInterceptors(
		tracingInterceptor,
		observability.LoggingInterceptor(logger),
		observability.RecoveryInterceptor(logger),
		observability.ValidationRequestInterceptor(),
		observability.ValidationRequestInterceptor(),
	)
	mux := http.NewServeMux()

	reflector := grpcreflect.NewStaticReflector(
		userconnect.UserCommandServiceName,
		otpconnect.OTPCommandServiceName,
	)

	userServer := user.NewUserServer(app.RegisterExecutor)
	otpServer := otp.NewOTPServer(app.VerifyExecutor, app.ResendExecutor)

	mux.Handle(userconnect.NewUserCommandServiceHandler(userServer, interceptors))
	mux.Handle(otpconnect.NewOTPCommandServiceHandler(otpServer, interceptors))

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return &AdapterModule{
		Mux: mux,
	}
}
