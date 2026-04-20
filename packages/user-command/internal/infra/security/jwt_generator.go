package security

import (
	"fmt"
	"shopify-user-command-module/internal/application/port"
	"time"

	"github.com/golang-jwt/jwt/v4"
	configx "github.com/iamKienb/shopify-go-platform/config"
)

type jwtClaims struct {
	UserID          string `json:"uid"`
	Email           string `json:"email"`
	PasswordVersion int    `json:"pwd_v"`
	jwt.RegisteredClaims
}

type JWTGenerator struct {
	cfg configx.JwtConfig
}

func NewJWTGenerator(cfg configx.JwtConfig) port.TokenGenerator {
	return &JWTGenerator{
		cfg: cfg,
	}
}

func (g *JWTGenerator) GeneratePair(claims port.TokenClaims) (*port.TokenPair, error) {
	now := time.Now().UTC()
	accessExpAt := now.Add(g.cfg.AccessExpiry)
	refreshExpAt := now.Add(g.cfg.RefreshExpiry)

	accessToken, err := g.sign(claims, accessExpAt, g.cfg.AccessSecret)
	if err != nil {
		return nil, fmt.Errorf("sign access token: %w", err)
	}

	refreshToken, err := g.sign(claims, refreshExpAt, g.cfg.RefreshSecret)
	if err != nil {
		return nil, fmt.Errorf("sign refresh token: %w", err)
	}

	return &port.TokenPair{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessExpAt,
		RefreshTokenExpiresAt: refreshExpAt,
	}, nil
}

func (g *JWTGenerator) sign(claims port.TokenClaims, expiryAt time.Time, secret string) (string, error) {
	claim := jwtClaims{
		UserID:          claims.UserId,
		Email:           claims.Email,
		PasswordVersion: claims.PasswordVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiryAt),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    "user-command",
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claim).SignedString([]byte(secret))

	if err != nil {
		return "", err
	}
	return token, nil
}
