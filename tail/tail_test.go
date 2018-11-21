package tail

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"context"
)

func TestReadTail(t *testing.T) {
	f, _ := os.Open("fixture/tail/test.log")
	defer f.Close()

	ctx, c := context.WithCancel(context.Background())
	lines, err := ReadTail(ctx, f)
	assert.Nil(t, err, "Error reading tail line")

	var result []string
	for {
		line, more := <-lines
		if more {
			result = append(result, line)
		} else {
			break
		}
	}

	assert.Equal(t, 5, len(result))
	assert.Equal(t, "", result[0])
	assert.Equal(t, "line3", result[1])
	assert.Equal(t, "", result[2])
	assert.Equal(t, "line2", result[3])
	assert.Equal(t, "line1", result[4])

	c()
	<- ctx.Done()
}
