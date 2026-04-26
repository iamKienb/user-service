package port

import (
	"shopify-user-command-module/internal/domain/account"
	"time"
)

type UserClaims struct {
	UserID          string
	Email           string
	Roles           []account.UserRole
	PasswordVersion int
}

type TokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

type TokenService interface {
	GeneratePair(claims UserClaims) (*TokenPair, error)
}
