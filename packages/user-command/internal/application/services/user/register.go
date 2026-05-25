package user

import (
	"context"
	"fmt"
	"time"

	"user-command-module/internal/application/commands/register_user"
	"user-command-module/internal/application/port"
	"user-command-module/internal/application/shared"
	"user-command-module/internal/domain/profile"
	domain_shared "user-command-module/internal/domain/shared"
	"user-command-module/internal/domain/user"

	"github.com/google/uuid"
)

func (s *userService) Register(ctx context.Context, cmd register_user.Command) (*register_user.Result, error) {
	if err := s.ensureEmailAvailable(ctx, cmd.Email); err != nil {
		return nil, s.wrapError(err)
	}

	passwordHash, err := s.hasher.Hash(cmd.Password)
	if err != nil {
		return nil, s.wrapError(err)
	}

	validGender := domain_shared.ValidateEnum[profile.GenderEnum](cmd.Profile.Gender)
	if validGender == nil {
		return nil, s.wrapError(profile.ErrGenderInvalid)
	}

	newUser := user.NewUser(user.NewUserParams{
		Email:        cmd.Email,
		PasswordHash: passwordHash,
	})

	newProfile := profile.NewProfile(profile.NewProfileParams{
		UserID:   newUser.ID,
		FullName: cmd.Profile.FullName,
		Gender:   *validGender,
	})

	if err := s.txManager.WithTx(ctx, func(ctx context.Context) error {
		if err := s.userRepo.CreateUser(ctx, newUser); err != nil {
			return err
		}

		if userEvents := newUser.FlushEvents(); len(userEvents) > 0 {
			if err := s.outboxService.Publish(ctx, port.OutboxParam{
				AggregateID:   newUser.ID.RawID(),
				AggregateType: newUser.Type(),
				Events:        userEvents,
			}); err != nil {
				return err
			}
		}

		if err := s.profileRepo.CreateProfile(ctx, newProfile); err != nil {
			return err
		}

		if profileEvents := newProfile.FlushEvents(); len(profileEvents) > 0 {
			if err := s.outboxService.Publish(ctx, port.OutboxParam{
				AggregateID:   newUser.ID.RawID(),
				AggregateType: newProfile.Type(),
				Events:        profileEvents,
			}); err != nil {
				return err
			}

		}

		return nil
	}); err != nil {
		return nil, s.wrapError(err)
	}
	_ = s.userCache.MarkEmailTaken(ctx, cmd.Email, shared.EmailCacheTTL)

	res, err := s.createRegistrationChallenge(ctx, newUser)
	if err != nil {
		return nil, s.wrapError(err)
	}

	return res, nil
}

func (s *userService) ensureEmailAvailable(ctx context.Context, email string) error {
	taken, _ := s.userCache.IsEmailTaken(ctx, email)
	if taken {
		return user.ErrEmailTaken
	}

	existing, err := s.userRepo.FindUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if existing != nil {
		_ = s.userCache.MarkEmailTaken(ctx, email, shared.EmailCacheTTL)
		return user.ErrEmailTaken
	}
	return nil
}

func (s *userService) createRegistrationChallenge(ctx context.Context, user *user.User) (*register_user.Result, error) {
	otp, err := shared.GenerateOTP()
	if err != nil {
		return nil, err
	}

	sessionToken := uuid.NewString()
	now := time.Now().UTC()

	if err := s.otpCache.SaveOTP(ctx, sessionToken, port.OTPEntry{
		OTP:       otp,
		UserID:    user.ID.String(),
		Email:     user.Email,
		ExpiresAt: now.Add(shared.OTPTTL),
	}, shared.OTPTTL); err != nil {
		return nil, err
	}

	if err := s.otpCache.SaveSession(ctx, sessionToken, port.SessionEntry{
		UserID: user.ID.String(),
		Email:  user.Email,
	}, shared.SessionTTL); err != nil {
		return nil, err
	}

	fmt.Printf("[DEBUG] SESSION: %s, OTP: %s\n", sessionToken, otp)

	return &register_user.Result{
		SessionToken: sessionToken,
		ExpiresAt:    now.Add(shared.SessionTTL),
	}, nil
}
