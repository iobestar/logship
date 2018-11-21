package unit

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"strings"
	"regexp"
	"time"
	"io/ioutil"
	"os"
	"path"
	"fmt"
)

func TestGetLogFilesInOrderOfModification(t *testing.T) {

	dir, _ := ioutil.TempDir("", "logshiptest")
	defer os.RemoveAll(dir)

	f1, err := os.Create(path.Join(dir, "test.1.log"))
	assert.Nil(t, err, "Error creating first file")
	err = os.Chtimes(f1.Name(), time.Now(), time.Now())
	f2, err := os.Create(path.Join(dir, "test.log"))
	assert.Nil(t, err, "Error creating second file")
	err = os.Chtimes(f2.Name(), time.Now(), time.Now().Add(1 * time.Second))

	defer f1.Close()
	defer f2.Close()

	fmt.Println(dir)
	unit := &LogUnit{
		Id:          "test",
		FilePattern: dir + "/test*.log",
	}

	logFiles, _ := unit.getLogFiles()

	assert.Equal(t, 2, len(logFiles))
	assert.Equal(t, true, strings.HasSuffix(logFiles[0].Path, "test.log"))
	assert.Equal(t, true, strings.HasSuffix(logFiles[1].Path, "test.1.log"))
}

func TestGetLogLimitedByCount(t *testing.T) {

	unit := &LogUnit{
		Id:             "test",
		FilePattern:    "fixture/output.log",
		LogPattern:     regexp.MustCompile(`^(?P<datetime>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3})\s+(?P<level>\w+).*`),
		DateTimeLayout: "2006-01-02 15:04:05.000",
	}

	logs := make([]*LogEntry, 0)
	unit.GetNLogs(2, func(logEntry *LogEntry) error {
		logs = append(logs, logEntry)
		return nil
	})

	assert.Equal(t, true, strings.HasPrefix(logs[0].Log(), "2018-09-06 22:58:15.434"))
	assert.Equal(t, true, strings.HasPrefix(logs[1].Log(), "2018-09-06 22:58:15.415"))
}

func TestGetLogsLimitedByTime(t *testing.T) {

	unit := &LogUnit{
		Id:             "test",
		FilePattern:    "fixture/output.log",
		LogPattern:     regexp.MustCompile(`^(?P<datetime>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3})\s+(?P<level>\w+).*`),
		DateTimeLayout: "2006-01-02 15:04:05.000",
	}

	logs := make([]*LogEntry, 0)

	d, _ := time.ParseDuration("1h")
	unit.GetTLogs(d, 1536274695434000000, func(logEntry *LogEntry) error {
		logs = append(logs, logEntry)
		return nil
	})

	assert.Equal(t, 2, len(logs))
}

func TestMultipleLogFiles(t *testing.T) {

	unit := &LogUnit{
		Id:             "test",
		FilePattern:    "fixture/part*",
		LogPattern:     regexp.MustCompile(`^(?P<datetime>\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\.\d{3})\s+(?P<level>\w+).*`),
		DateTimeLayout: "2006-01-02 15:04:05.000",
	}

	var c = 0
	unit.getLogs(func(count int, lastEntry *LogEntry) bool {
		return false
	}, func(logEntry *LogEntry) error {
		c = c + 1
		return nil
	})

	assert.Equal(t, 5, c, "Incorrect number of logs")
}

func TestGetLinesLimitedByCount(t *testing.T) {

	unit := &LogUnit{
		Id:          "test",
		FilePattern: "fixture/lines.log",
	}

	var lines []string
	unit.GetNLines(3, func(line string) error {
		lines = append(lines, line)
		return nil
	})

	assert.EqualValues(t, 3, len(lines))
	assert.EqualValues(t, "", lines[0])
	assert.EqualValues(t, "third line", lines[1])
	assert.EqualValues(t, "", lines[2])
}
