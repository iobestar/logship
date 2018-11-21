package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLines(t *testing.T) {

	unit := &LogUnit{
		Id:          "test",
		FilePattern: "fixture/lines.log",
	}

	lines := make([]string, 0, 0)
	err := unit.getLines(func(count int, line string) (bool, error) {

		lines = append(lines, line)
		return count == 4, nil
	})

	assert.Nil(t, err, "get lines error not nil")
	assert.EqualValues(t, 4, len(lines))
	assert.EqualValues(t, "", lines[0])
	assert.EqualValues(t, "third line", lines[1])
	assert.EqualValues(t, "", lines[2])
	assert.EqualValues(t, "second line", lines[3])
}
