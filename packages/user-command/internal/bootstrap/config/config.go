package config

import (
	configx "github.com/iamKienb/shopify-go-platform/config"
)

type UserCommandConfig struct {
	Postgres configx.PostgresConfig `envPrefix:"USER_COMMAND_SERVICE"`
	Redis    configx.RedisConfig    `envPrefix:"USER_COMMAND_SERVICE"`
	Jwt      configx.JwtConfig      `envPrefix:"USER_COMMAND_SERVICE"`
	Argon2   configx.Argon2Config   `envPrefix:"USER_COMMAND_SERVICE"`
	Server   configx.Server         `envPrefix:"USER_COMMAND_SERVICE"`
}
