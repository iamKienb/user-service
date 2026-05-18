package shop

import (
	"context"

	"user-command-module/internal/application/command/create_shop"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/shared"

	domain_shared "user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/shop"
)

func (s *shopService) CreateShop(ctx context.Context, cmd create_shop.Command) (*create_shop.Result, error) {
	if err := s.checkSlugAvailable(ctx, cmd.Slug); err != nil {
		return nil, s.wrapError(err)
	}

	if err := s.checkIdempotency(ctx, cmd.User.ID); err != nil {
		return nil, s.wrapError(err)
	}

	profile := profileFromCommand(cmd.Profile)

	shopAgg := shop.NewAggregate(shop.AggregateParams{
		UserID:      cmd.User.ID,
		Name:        cmd.Name,
		Slug:        cmd.Slug,
		Description: profile.Description,
		LogoUrl:     profile.LogoUrl,
		BannerUrl:   profile.BannerUrl,
	})

	memberAgg := shop.NewMemberAggregate(shop.MemberAggregateParams{
		ShopID:      shopAgg.Shop.ID,
		MemberID:    cmd.User.ID,
		MemberName:  cmd.User.Name,
		AddedBy:     cmd.User.ID,
		NameAddedBy: cmd.User.Name,
		RoleIDs:     []domain_shared.RoleID{shop.RoleOwnerID},
	})

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.shopRepo.SaveAggregate(ctx, shopAgg); err != nil {
			return err
		}

		membersToSave := []*shop.MemberAggregate{memberAgg}
		if err := s.shopRepo.UpsertMemberAggregate(ctx, membersToSave); err != nil {
			return err
		}

		shopEvents := shopAgg.FlushEvents()
		if len(shopEvents) == 0 {
			return nil
		}

		outboxShopParam := port.OutboxParam{
			AggregateID:   shopAgg.Shop.ID.RawID(),
			AggregateType: shop.AggregateTypeShop,
			Events:        shopEvents,
		}
		if err := s.outboxService.Publish(ctx, outboxShopParam); err != nil {
			return err
		}

		memberEvents := memberAgg.FlushEvents()
		if len(memberEvents) == 0 {
			return nil
		}

		outboxMemberParam := port.OutboxParam{
			AggregateID:   memberAgg.Member.ShopID.RawID(),
			AggregateType: shop.AggregateTypeShop,
			Events:        memberEvents,
		}
		if err := s.outboxService.Publish(ctx, outboxMemberParam); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, s.wrapError(err)
	}

	bgCtx := context.WithoutCancel(ctx)
	go func() {
		_ = s.shopCache.SetIdemKey(bgCtx, cmd.User.ID, shared.IdemKeyTTL)
		_ = s.shopCache.AddSlugToBloomFilter(bgCtx, cmd.Slug)
	}()

	return &create_shop.Result{
		ShopID: shopAgg.Shop.ID.String(),
	}, nil
}

func (s *shopService) checkIdempotency(ctx context.Context, userID domain_shared.UserID) error {
	exists, err := s.shopCache.IsIdemKeyTaken(ctx, userID)
	if err != nil {
		return err
	}

	if exists {
		return shop.ErrShopConflict
	}

	return nil
}

func (s *shopService) checkSlugAvailable(ctx context.Context, slug string) error {
	exists, err := s.shopCache.GetSlugFromBloomFilter(ctx, slug)
	if err != nil {
		return err
	}

	if exists == 0 {
		return nil
	}

	isDuplicateSlug, err := s.shopRepo.CheckSlugExists(ctx, slug)
	if err != nil {
		return err
	}
	if isDuplicateSlug {
		return shop.ErrShopSlugTaken
	}

	return nil
}

func profileFromCommand(profile *create_shop.Profile) create_shop.Profile {
	if profile == nil {
		return create_shop.Profile{}
	}

	return *profile
}
