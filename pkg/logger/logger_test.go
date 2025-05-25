package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewLogger(t *testing.T) {
	log, err := NewLogger()
	assert.NoError(t, err)
	assert.NotNil(t, log)
}

func TestZapLogger_InfoAndError(t *testing.T) {
	log, _ := NewLogger()
	defer log.Sync()

	// Basic smoke test - should not panic
	log.Info("info message")
	log.Error("error message", zap.String("key", "value"))
}
