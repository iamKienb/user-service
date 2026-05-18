package module

import (
	"user-command-module/internal/application/command/add_shop_address"
	"user-command-module/internal/application/command/add_user_address"
	"user-command-module/internal/application/command/assign_member"
	"user-command-module/internal/application/command/create_shop"
	"user-command-module/internal/application/command/login_user"
	"user-command-module/internal/application/command/register_user"
	"user-command-module/internal/application/command/resend_otp"
	"user-command-module/internal/application/command/verify_otp"
	"user-command-module/internal/application/command/verify_permission"
	"user-command-module/internal/application/service/otp"
	"user-command-module/internal/application/service/outbox"
	"user-command-module/internal/application/service/shop"
	"user-command-module/internal/application/service/user"
	"user-command-module/internal/domain/auth"
)

type ApplicationModule struct {
	RegisterExecutor       register_user.Executor
	LoginExecutor          login_user.Executor
	AddUserAddressExecutor add_user_address.Executor

	VerifyExecutor verify_otp.Executor
	ResendExecutor resend_otp.Executor

	CreateShopExecutor       create_shop.Executor
	AssignMemberExecutor     assign_member.Executor
	AddShopAddressExecutor   add_shop_address.Executor
	VerifyPermissionExecutor verify_permission.Executor
}

func NewApplicationModule(infra *InfraModule) *ApplicationModule {
	outboxService := outbox.NewOutboxService(infra.OutboxRepo)
	authorizer := auth.NewAuthorizer()

	userService := user.NewUserService(
		infra.AccountRepo,
		infra.AuthRepo,
		outboxService,

		infra.UserCache,
		infra.OtpCache,

		infra.TokenGenerator,
		infra.TxManager,
		infra.Hasher,
	)

	otpService := otp.NewOTPService(
		infra.AccountRepo,
		outboxService,

		infra.TokenGenerator,
		infra.OtpCache,
		infra.TxManager,
	)

	shopService := shop.NewShopService(
		infra.ShopRepo,
		authorizer,
		outboxService,
		infra.ShopCache,
		infra.TxManager,
	)

	registerCommandHandler := register_user.NewHandler(userService)
	loginCommandHandler := login_user.NewHandler(userService)
	addUserAddressHandler := add_user_address.NewHandler(userService)

	verifyCommandHandler := verify_otp.NewHandler(otpService)
	resendCommandHandler := resend_otp.NewHandler(otpService)

	createShopCommandHandler := create_shop.NewHandler(shopService)
	assignMemberHandler := assign_member.NewHandler(shopService)
	addShopAddressHandler := add_shop_address.NewHandler(shopService)
	verifyPermissionCommandHandler := verify_permission.NewHandler(shopService)

	return &ApplicationModule{
		RegisterExecutor:       registerCommandHandler,
		LoginExecutor:          loginCommandHandler,
		AddUserAddressExecutor: addUserAddressHandler,

		VerifyExecutor: verifyCommandHandler,
		ResendExecutor: resendCommandHandler,

		CreateShopExecutor:       createShopCommandHandler,
		AssignMemberExecutor:     assignMemberHandler,
		AddShopAddressExecutor:   addShopAddressHandler,
		VerifyPermissionExecutor: verifyPermissionCommandHandler,
	}
}
