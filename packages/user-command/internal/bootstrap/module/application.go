package module

import (
	"user-command-module/internal/application/commands/add_user_address"
	"user-command-module/internal/application/commands/login_user"
	"user-command-module/internal/application/commands/register_user"
	"user-command-module/internal/application/commands/resend_otp"
	"user-command-module/internal/application/commands/verify_otp"
	"user-command-module/internal/application/services/otp"
	"user-command-module/internal/application/services/outbox"
	"user-command-module/internal/application/services/user"
)

type ApplicationModule struct {
	RegisterExecutor       register_user.Executor
	LoginExecutor          login_user.Executor
	AddUserAddressExecutor add_user_address.Executor

	VerifyExecutor verify_otp.Executor
	ResendExecutor resend_otp.Executor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	outboxService := outbox.NewOutboxService(infra.OutboxRepo)

	userService := user.NewUserService(
		infra.UserRepo,
		infra.AuthRepo,
		infra.ProfileRepo,
		infra.UserAddressRepo,
		outboxService,

		infra.UserCache,
		infra.OtpCache,

		infra.TokenGenerator,
		infra.TxManager,
		infra.Hasher,
	)

	otpService := otp.NewOTPService(
		infra.UserRepo,
		infra.ProfileRepo,
		outboxService,

		infra.TokenGenerator,
		infra.OtpCache,
		infra.TxManager,
	)

	registerCommandHandler := register_user.NewHandler(userService)
	loginCommandHandler := login_user.NewHandler(userService)
	addUserAddressHandler := add_user_address.NewHandler(userService)

	verifyCommandHandler := verify_otp.NewHandler(otpService)
	resendCommandHandler := resend_otp.NewHandler(otpService)

	return &ApplicationModule{
		RegisterExecutor:       registerCommandHandler,
		LoginExecutor:          loginCommandHandler,
		AddUserAddressExecutor: addUserAddressHandler,

		VerifyExecutor: verifyCommandHandler,
		ResendExecutor: resendCommandHandler,
	}
}
