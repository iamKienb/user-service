package port

import (
	"time"
)

type TokenClaims struct {
	UserId          string
	Email           string
	PasswordVersion int
}

type TokenPair struct {
	AccessToken           string
	RefreshToken          string
	AccessTokenExpiresAt  time.Time
	RefreshTokenExpiresAt time.Time
}

type TokenGenerator interface {
	GeneratePair(claims TokenClaims) (*TokenPair, error)
}
