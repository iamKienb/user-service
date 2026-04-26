package security

import (
	"shopify-user-command-module/internal/application/port"
	"shopify-user-command-module/internal/infra/common"

	configx "github.com/iamKienb/shopify-go-platform/config"
	authx "github.com/iamKienb/shopify-go-platform/middleware/auth"
)

type TokenGenerator struct {
	authx.Generator
}

func NewTokenGenerator(cfg configx.JwtConfig) port.TokenService {
	return &TokenGenerator{
		Generator: authx.NewJWTGenerator(cfg),
	}
}

func (g *TokenGenerator) GeneratePair(claims port.UserClaims) (*port.TokenPair, error) {
	pair, err := g.Generator.GeneratePair(authx.Claims{
		UserID:          claims.UserID,
		Email:           claims.Email,
		Roles:           common.ToStringRoles(claims.Roles),
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
