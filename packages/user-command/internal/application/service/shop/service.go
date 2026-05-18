package shop

import (
	"context"
	"user-command-module/internal/application/command/add_shop_address"
	"user-command-module/internal/application/command/assign_member"
	"user-command-module/internal/application/command/create_shop"
	"user-command-module/internal/application/command/verify_permission"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/service/outbox"
	"user-command-module/internal/domain/auth"
	"user-command-module/internal/domain/shop"
)

type Service interface {
	CreateShop(ctx context.Context, cmd create_shop.Command) (*create_shop.Result, error)
	AssignMember(ctx context.Context, cmd assign_member.Command) (*assign_member.Result, error)
	AddAddress(ctx context.Context, cmd add_shop_address.Command) (*add_shop_address.Result, error)
	VerifyPermission(ctx context.Context, cmd verify_permission.Command) (*verify_permission.Result, error)
}

type shopService struct {
	shopRepo      shop.Repository
	authorizer    auth.IAuthorizer
	outboxService outbox.Service
	shopCache     port.ShopCache
	txManager     port.TxManager
}

func NewShopService(
	shopRepo shop.Repository,
	authorizer auth.IAuthorizer,
	outboxService outbox.Service,
	shopCache port.ShopCache,
	txManager port.TxManager,
) Service {
	return &shopService{
		shopRepo:      shopRepo,
		authorizer:    authorizer,
		outboxService: outboxService,
		shopCache:     shopCache,
		txManager:     txManager,
	}
}
