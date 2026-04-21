package module

import (
	"shopify-user-command-module/internal/application/command/login_user"
	"shopify-user-command-module/internal/application/command/register_user"
	"shopify-user-command-module/internal/application/command/resend_otp"
	"shopify-user-command-module/internal/application/command/verify_otp"
	"shopify-user-command-module/internal/application/service/otp"
	"shopify-user-command-module/internal/application/service/user"
)

type ApplicationModule struct {
	RegisterExecutor register_user.Executor
	LoginExecutor    login_user.Executor
	VerifyExecutor   verify_otp.Executor
	ResendExecutor   resend_otp.Executor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	userService := user.NewUserService(
		infra.IdentityRepo,
		infra.Cache,
		infra.OtpCache,
		infra.TokenGenerator,
		infra.TxManager,
		infra.Hasher,
	)

	otpService := otp.NewOTPService(
		infra.IdentityRepo,
		infra.TokenGenerator,
		infra.OtpCache,
		infra.TxManager,
	)

	registerCommandHandler := register_user.NewHandler(userService)
	loginCommandHandler := login_user.NewHandler(userService)

	verifyCommandHandler := verify_otp.NewHandler(otpService)
	resendCommandHandler := resend_otp.NewHandler(otpService)

	return &ApplicationModule{
		RegisterExecutor: registerCommandHandler,
		LoginExecutor:    loginCommandHandler,
		VerifyExecutor:   verifyCommandHandler,
		ResendExecutor:   resendCommandHandler,
	}
}
