package config

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewConfigWithDefaults(t *testing.T) {

	filename := "fixture/logunit.yml"

	c, err := Parse(filename)
	assert.Nil(t, err)

	assert.Equal(t, 2, len(c.LogUnits))

	assert.Equal(t, "rs_replication_log", c.LogUnits[0].Id)
	assert.Equal(t, "replication.log", c.LogUnits[0].FilePattern)
	assert.Equal(t, defaultLogPattern, c.LogUnits[0].LogPattern)
	assert.Equal(t, defaultDateTimeLayout, c.LogUnits[0].DateTimeLayout)

	assert.Equal(t, "rs_error_log", c.LogUnits[1].Id)
	assert.Equal(t, "error.log", c.LogUnits[1].FilePattern)
	assert.Equal(t, "any", c.LogUnits[1].LogPattern)
	assert.Equal(t, "42", c.LogUnits[1].DateTimeLayout)
}
