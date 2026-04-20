package cache

import (
	"context"
	"fmt"
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/domain/identity"
	"time"

	redisx "github.com/iamKienb/shopify-go-platform/redis"
	"github.com/redis/go-redis/v9"
)

type otpCache struct {
	redisx.Cache
	client *redis.Client
}

func NewOTPCache(client *redis.Client) port.OTPCache {
	return &otpCache{
		Cache:  redisx.NewRedisService(client),
		client: client,
	}
}

const (
	otpKey      = "user-command:otp:%s"
	sessionKey  = "user-command:session:%s"
	resendKey   = "user-command:resend:%s"
	KeyNotFound = "key not found"
)

func (c *otpCache) SaveOTP(ctx context.Context, sessionToken string, otp port.OTPEntry, ttl time.Duration) error {
	key := fmt.Sprintf(otpKey, sessionToken)
	fmt.Println("SAVE OTP", key)
	if err := c.Set(ctx, key, otp, ttl); err != nil {
		return fmt.Errorf("redis: set otp: %w", err)
	}
	return nil
}

func (c *otpCache) GetOTP(ctx context.Context, sessionToken string) (*port.OTPEntry, error) {
	key := fmt.Sprintf(otpKey, sessionToken)
	var otp port.OTPEntry

	if err := c.Get(ctx, key, &otp); err != nil {
		if err == redis.Nil {
			return nil, identity.ErrOTPInvalid
		}

		return nil, fmt.Errorf("redis: get otp: %w", err)
	}

	return &otp, nil
}
func (c *otpCache) DeleteOTP(ctx context.Context, sessionToken string) error {
	key := fmt.Sprintf(otpKey, sessionToken)
	if err := c.Delete(ctx, key); err != nil {
		return fmt.Errorf("redis: delete otp: %w", err)
	}

	return nil
}

func (c *otpCache) SaveSession(ctx context.Context, sessionToken string, session port.SessionEntry, ttl time.Duration) error {
	key := fmt.Sprintf(sessionKey, sessionToken)
	if err := c.Set(ctx, key, session, ttl); err != nil {
		return fmt.Errorf("redis: set session: %w", err)
	}

	return nil
}

func (c *otpCache) GetSession(ctx context.Context, sessionToken string) (*port.SessionEntry, error) {
	key := fmt.Sprintf(sessionKey, sessionToken)
	var session port.SessionEntry

	err := c.Get(ctx, key, &session)
	if err != nil {
		if err == redis.Nil {
			return nil, identity.ErrSessionInvalid
		}

		return nil, fmt.Errorf("redis: get session: %w", err)
	}

	return &session, nil
}

func (c *otpCache) DeleteSession(ctx context.Context, sessionToken string) error {
	key := fmt.Sprintf(sessionKey, sessionToken)
	if err := c.Delete(ctx, key); err != nil {
		return fmt.Errorf("redis: delete session: %w", err)
	}

	return nil
}

func (c *otpCache) IncrResendCount(ctx context.Context, sessionToken string, ttl time.Duration) (int64, error) {
	key := fmt.Sprintf(resendKey, sessionToken)

	count, err := c.client.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis: incr resend count: %w", err)
	}

	if count == 1 {
		_ = c.client.Expire(ctx, key, ttl)
	}

	return count, nil
}
