package module

import (
	"log/slog"
	"net/http"

	"user-command-module/internal/adapter/otp"
	"user-command-module/internal/adapter/user"

	"connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"github.com/iamKienb/api-contract/gen/otp/otpconnect"
	"github.com/iamKienb/api-contract/gen/user/userconnect"
	authx "github.com/iamKienb/go-core/middleware/auth"
	observabilityx "github.com/iamKienb/go-core/middleware/observability"
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
		authx.RequestContextInterceptor(),
		authx.AuthInternalInterceptor(),
		observabilityx.LoggingInterceptor(logger),
		observabilityx.ValidationRequestInterceptor(),
		observabilityx.ErrorResponseInterceptor(logger),
	)

	allInterceptors := connect.WithInterceptors(interceptors...)

	mux := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		userconnect.UserCommandServiceName,
		otpconnect.OTPCommandServiceName,
	)

	userServer := user.NewUserServer(
		app.RegisterExecutor,
		app.LoginExecutor,
		app.AddUserAddressExecutor,
	)

	otpServer := otp.NewOTPServer(
		app.VerifyExecutor,
		app.ResendExecutor,
	)

	mux.Handle(userconnect.NewUserCommandServiceHandler(userServer, allInterceptors))
	mux.Handle(otpconnect.NewOTPCommandServiceHandler(otpServer, allInterceptors))

	mux.Handle(grpcreflect.NewHandlerV1(reflector))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))

	return &AdapterModule{Mux: mux}
}
