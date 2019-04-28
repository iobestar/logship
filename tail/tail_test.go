package tail

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadTail(t *testing.T) {
	f, _ := os.Open("fixture/test.log")
	defer f.Close()

	ctx, cancel := context.WithCancel(context.Background())
	lines, errors := ReadTail(ctx, f)

	var result []string
	done:
	for {
		select {
		case line, ok := <-lines:
			if ok {
				result = append(result, line)
			} else {
				break done
			}
		case err, ok := <-errors:
			if ok {
				assert.Nil(t, err, "unexpected error during tails reading")
			} else {
				break done
			}
		}
	}

	assert.Equal(t, 4, len(result))
	assert.Equal(t, "line3", result[0])
	assert.Equal(t, "", result[1])
	assert.Equal(t, "line2", result[2])
	assert.Equal(t, "line1", result[3])

	cancel()
	<- ctx.Done()
}
