package unit

import (
	"context"
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestLogUnit_StreamLogs(t *testing.T) {

	logUnit := &LogUnit{
		Id:   "test",
		Glob: "fixture/output.log",
	}

	dateTimeLayout := "2006-01-02 15:04:05.000"
	logPattern := regexp.MustCompile(`^(?P<datetime>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3})\s+(?P<level>\w+).*`)
	ctx, cancel := context.WithCancel(context.Background())
	logs, errors := logUnit.StreamLogs(ctx, dateTimeLayout, logPattern)

	var result []*LogEntry
done:
	for {
		select {
		case log, ok := <-logs:
			if ok {
				result = append(result, log)
			} else {
				break done
			}
		case err, ok := <-errors:
			if ok {
				assert.Nil(t, err, "unexpected error during log reading")
			}
		}
	}
	cancel()
	<-ctx.Done()

	assert.Equal(t, 5, len(result))

	assert.True(t, strings.HasPrefix(result[0].Log(), "2018-09-06 22:58:15.434"))
	assert.True(t, strings.HasPrefix(result[4].Log(), "2018-08-08 18:58:15.311"))
}
