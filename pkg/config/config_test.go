package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig_Defaults(t *testing.T) {
	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "8080", cfg.ServerPort)
	assert.Equal(t, "mongodb://localhost:27017", cfg.MongoURI)
	assert.Equal(t, 10*time.Second, cfg.MongoTimeout)
	assert.Equal(t, "secret", cfg.JWTSecret)
	assert.Equal(t, 60, cfg.RateLimitPerMin)
}

func TestLoadConfig_EnvOverride(t *testing.T) {
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("MONGO_URI", "mongodb://mongo:27017")
	os.Setenv("MONGO_TIMEOUT", "5")
	os.Setenv("JWT_SECRET", "supersecret")
	os.Setenv("RATE_LIMIT_PER_MIN", "100")

	defer func() {
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("MONGO_URI")
		os.Unsetenv("MONGO_TIMEOUT")
		os.Unsetenv("JWT_SECRET")
		os.Unsetenv("RATE_LIMIT_PER_MIN")
	}()

	cfg, err := LoadConfig()
	assert.NoError(t, err)
	assert.Equal(t, "9090", cfg.ServerPort)
	assert.Equal(t, "mongodb://mongo:27017", cfg.MongoURI)
	assert.Equal(t, 5*time.Second, cfg.MongoTimeout)
	assert.Equal(t, "supersecret", cfg.JWTSecret)
	assert.Equal(t, 100, cfg.RateLimitPerMin)
}
