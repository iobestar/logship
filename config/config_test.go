package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig("fixture/test-config.yml")
	assert.Nil(t, err)

	assert.Equal(t, 2, len(cfg.LogUnits))

	assert.Equal(t, "test_error", cfg.LogUnits[0].Id)
	assert.Equal(t, "error.*", cfg.LogUnits[0].Glob)

	assert.Equal(t, "test_output", cfg.LogUnits[1].Id)
	assert.Equal(t, "output.*", cfg.LogUnits[1].Glob)
}
