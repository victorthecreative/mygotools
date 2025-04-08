package config

import (
	cerrs "github.com/cockroachdb/errors"
	"github.com/spf13/viper"
)

// Configuration: project settings.
type Configuration struct {
	Port string
	// Configurações para PostgreSQL
	PostgresUser     string
	PostgresPassword string
	PostgresHost     string
	PostgresDriver   string
	PostgresDatabase string
	PostgresPort     string

	// Configurações para MongoDB
	MongoURI            string
	MongoDatabase       string
	MongoUsername       string
	MongoPassword       string
	MongoAuthSource     string
	MongoConnectTimeout int

	// Configurações para Redis
	RedisHost       string
	RedisPort       string
	RedisPassword   string
	RedisDB         int
	RedisMaxRetries int
	RedisTimeout    int
	RedisPoolSize   int

	AcceptConfigurationEvents bool
	AcceptDiscoveryEvents     bool
}

// New load all settings for this service.
func NewConfig(typeConfig string) (*Configuration, error) {
	v := viper.New()

	v.AutomaticEnv()
	if typeConfig == "env" {
		v.SetConfigName(".env")
		v.SetConfigType("env")
	} else if typeConfig == "yaml" {
		v.SetConfigName("config_yaml")
		v.SetConfigType("yaml")
	} else if typeConfig == "json" {
		v.SetConfigName("config_json")
		v.SetConfigType("json")
	}
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		return nil, cerrs.WithStack(err)
	}

	var cfg Configuration
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, cerrs.WithStack(err)
	}

	return &cfg, nil
}
