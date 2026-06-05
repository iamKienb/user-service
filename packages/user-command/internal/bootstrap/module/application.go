package module

import (
	"user-command-module/internal/application/commands/add_user_address"
	"user-command-module/internal/application/commands/login_user"
	"user-command-module/internal/application/commands/register_user"
	"user-command-module/internal/application/commands/resend_otp"
	"user-command-module/internal/application/commands/verify_otp"
	get_user_address_by_id "user-command-module/internal/application/queries/get_address_by_id"
	"user-command-module/internal/application/services/otp"
	"user-command-module/internal/application/services/outbox"
	"user-command-module/internal/application/services/user"
)

type ApplicationModule struct {
	RegisterExecutor       register_user.Executor
	LoginExecutor          login_user.Executor
	AddUserAddressExecutor add_user_address.Executor
	GetUserAddressExecutor get_user_address_by_id.Executor

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

	return &ApplicationModule{
		RegisterExecutor:       register_user.NewHandler(userService),
		LoginExecutor:          login_user.NewHandler(userService),
		AddUserAddressExecutor: add_user_address.NewHandler(userService),
		GetUserAddressExecutor: get_user_address_by_id.NewHandler(userService),

		VerifyExecutor: verify_otp.NewHandler(otpService),
		ResendExecutor: resend_otp.NewHandler(otpService),
	}
}
