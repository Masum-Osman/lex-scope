package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort      string        `mapstructure:"SERVER_PORT"`
	MongoURI        string        `mapstructure:"MONGO_URI"`
	MongoTimeout    time.Duration `mapstructure:"MONGO_TIMEOUT"`
	RedisAddr       string        `mapstructure:"REDIS_ADDR"`
	RedisPassword   string        `mapstructure:"REDIS_PASSWORD"`
	RedisDB         int           `mapstructure:"REDIS_DB"`
	JWTSecret       string        `mapstructure:"JWT_SECRET"`
	RateLimitPerMin int           `mapstructure:"RATE_LIMIT_PER_MIN"`
}

func LoadConfig() (*Config, error) {
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("MONGO_URI", "mongodb://localhost:27017")
	viper.SetDefault("MONGO_TIMEOUT", 10)
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DB", 0)
	viper.SetDefault("JWT_SECRET", "secret")
	viper.SetDefault("RATE_LIMIT_PER_MIN", 60)

	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	cfg.MongoTimeout = cfg.MongoTimeout * time.Second

	return &cfg, nil
}
