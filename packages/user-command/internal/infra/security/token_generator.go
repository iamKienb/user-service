package security

import (
	"user-command-module/internal/application/port"
	"user-command-module/internal/domain/shared"

	jwtx "github.com/iamKienb/go-core/jwt"
)

type TokenGenerator struct {
	service jwtx.JWTXService
}

func NewTokenGenerator(service jwtx.JWTXService) port.TokenService {
	return &TokenGenerator{
		service: service,
	}
}

func (g *TokenGenerator) GeneratePair(claims port.UserClaims) (*port.TokenPair, error) {
	pair, err := g.service.GeneratePair(jwtx.Claims{
		UserID:          claims.UserID,
		Email:           claims.Email,
		Roles:           shared.Strings(claims.Roles),
		PasswordVersion: claims.PasswordVersion,
	})
	if err != nil {
		return nil, err
	}

	return &port.TokenPair{
		AccessToken:           pair.AccessToken,
		RefreshToken:          pair.RefreshToken,
		AccessTokenExpiresAt:  pair.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: pair.RefreshTokenExpiresAt,
	}, nil
}
