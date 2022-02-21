package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfig(t *testing.T) {
	var cfg *Config

	t.Run("same-type", func(t *testing.T) {
		assert.IsType(t, cfg, instance)
	})
}
