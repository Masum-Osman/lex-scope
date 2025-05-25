package db

import (
	"testing"

	"github.com/Masum-Osman/lex-scope/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestNewMongoClient(t *testing.T) {
	cfg, err := config.LoadConfig()
	assert.NoError(t, err)

	client, err := NewMongoClient(cfg)
	if err != nil {
		t.Skip("MongoDB not available, skipping test")
		return
	}

	assert.NotNil(t, client)
}
