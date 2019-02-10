package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfigWithDefaults(t *testing.T) {

	filename := "fixture/test-config.yml"

	c, err := ParseConfig(filename)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(c.LogReaders))

	assert.Equal(t, "rs_replication_log", c.LogReaders[0].Id)
	assert.Equal(t, defaultLogPattern, c.LogReaders[0].LogPattern)
	assert.Equal(t, defaultDateTimeLayout, c.LogReaders[0].DateTimeLayout)

	assert.Equal(t, "rs_error_log", c.LogReaders[1].Id)
	assert.Equal(t, "any", c.LogReaders[1].LogPattern)
	assert.Equal(t, "42", c.LogReaders[1].DateTimeLayout)
}
