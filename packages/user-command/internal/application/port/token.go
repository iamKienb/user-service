package port

import (
	"time"
	"user-command-module/internal/domain/account"
)

type UserClaims struct {
	UserID          string
	Email           string
	FullName        string
	Roles           []account.RoleEnum
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
