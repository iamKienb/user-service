package config

import configx "github.com/iamKienb/go-core/config"

type UserQueryConfig struct {
	ES configx.ElasticSearchConfig `envPrefix:"USER_QUERY_SERVICE"`
}
