package config

import (
	"dogchecker/pkg/mailer"
	"dogchecker/pkg/petfinder"
	"os"
)

const (
	EnvProduction  = "production"
	EnvDevelopment = "development"
)

type Config struct {
	Environment string
	Petfinder   *petfinder.Config
	Redis       *RedisConfig
	Mailer      *mailer.Config
}

type RedisConfig struct {
	BaseUrl string
}

type MailerConfig struct {
	ApiKey      string
	FromAddress string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		Environment: getEnv("ENVIRONMENT", EnvDevelopment),
		Petfinder: &petfinder.Config{
			BaseUrl:      getEnv("PETFINDER_BASE_URL", "https://api.petfinder.com/v2"),
			ClientId:     getEnv("PETFINDER_CLIENT_ID", "apiKey"),
			ClientSecret: getEnv("PETFINDER_CLIENT_SECRET", "apiSecret"),
		},
		Redis: &RedisConfig{
			BaseUrl: getEnv("REDIS_URL", "redis:6379"),
		},
		Mailer: &mailer.Config{
			ApiKey:      getEnv("MAILER_API_KEY", "apiKey"),
			FromName:    getEnv("MAILER_FROM_NAME", "Michael Scott"),
			FromAddress: getEnv("MAILER_FROM_ADDRESS", "mscott@dundermifflin.com"),
			ToName:      getEnv("MAILER_TO_NAME", "Todd Packer"),
			ToAddress:   getEnv("MAILER_TO_ADDRESS", "packaging@dundermifflin.com"),
		},
	}

	return config, nil
}

func getEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}
