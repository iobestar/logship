package unit

import (
	"path/filepath"
	"os"
	"regexp"
	"time"
	"sort"
	"bytes"
	"strings"
	"github.com/iobestar/logship/config"
)

func NewLogUnit(cLogUnit config.LogUnitConfig) (*LogUnit, error) {

	lpRegex, err := regexp.Compile(cLogUnit.LogPattern)
	if nil != err {
		return nil, err
	}

	return &LogUnit{
		Id:             cLogUnit.Id,
		FilePattern:    cLogUnit.FilePattern,
		LogPattern:     lpRegex,
		DateTimeLayout: cLogUnit.DateTimeLayout,
	}, nil
}

type LogEntry struct {
	Timestamp int64
	Level     string
	payload   []string
}

func (entry *LogEntry) Log() string {

	var buffer bytes.Buffer
	for i := len(entry.payload) - 1; i >= 0; i-- {

		var line = strings.TrimSuffix(entry.payload[i], "\n")
		if len(line) == 0 && i == 0 {
			continue
		}

		if buffer.Len() > 0 {
			buffer.WriteString("\n" + line)
		} else {
			buffer.WriteString(line)
		}
	}
	return buffer.String()
}

type LogUnit struct {
	Id             string
	FilePattern    string
	LogPattern     *regexp.Regexp
	DateTimeLayout string
}

func (lu *LogUnit) GetNLines(n int, consume func(line string) (error)) error {

	if n == 0 {
		return nil
	}
	lu.getLines(func(count int, line string) (bool, error) {
		err := consume(line)
		if nil != err {
			return true, err
		}
		return count == n, nil
	})
	return nil
}

func (lu *LogUnit) GetNLogs(n int, consume func(logEntry *LogEntry) (error)) error {
	return lu.getLogs(func(count int, lastEntry *LogEntry) bool {
		return n == count-1
	}, consume)
}

func (lu *LogUnit) GetTLogs(t time.Duration, offset int64, consume func(logEntry *LogEntry) (error)) error {

	var lower = offset - t.Nanoseconds()
	return lu.getLogs(func(count int, lastEntry *LogEntry) bool {
		return lastEntry.Timestamp < lower
	}, consume)
}

func (lu *LogUnit) getLogs(isEnough func(count int, lastEntry *LogEntry) (bool), consume func(logEntry *LogEntry) (error)) error {

	var logPayload []string
	logCount := 0
	return lu.getLines(func(count int, line string) (bool, error) {
		logPayload = append(logPayload, line)
		match := lu.LogPattern.FindStringSubmatch(line)
		if len(match) > 0 && len(logPayload) > 0 {
			logEntry, err := lu.createLogEntry(lu.getLogParameters(match), logPayload)
			if nil != err {
				return true, err
			}
			logPayload = make([]string, 0)
			logCount = logCount + 1

			if isEnough(logCount, logEntry) {
				return true, nil
			}

			err = consume(logEntry)
			if nil != err {
				return true, err
			}
		}
		return false, nil
	})
}

func (lu *LogUnit) createLogEntry(params map[string]string, payload []string) (*LogEntry, error) {

	var dateTime string
	var timestamp int64
	var err error
	if value, valueExist := params["datetime"]; valueExist {
		dateTime = value
		timestamp, err = toTimestamp(lu.DateTimeLayout, dateTime)
		if nil != err {
			return nil, err
		}
	}

	var level string
	if value, valueExist := params["level"]; valueExist {
		level = value
	}

	return &LogEntry{
		Timestamp: timestamp,
		Level:     level,
		payload:   payload,
	}, nil

}

func toTimestamp(layout, dateTime string) (int64, error) {

	if len(dateTime) == 0 {
		return 0, nil
	}

	tm, err := time.Parse(layout, dateTime)
	if nil != err {
		return 0, err
	}
	return tm.UnixNano(), nil
}

func (lu *LogUnit) getLogParameters(match []string) map[string]string {

	result := make(map[string]string)
	for i, name := range lu.LogPattern.SubexpNames() {
		if i > 0 && i <= len(match) {
			result[name] = match[i]
		}
	}
	return result
}

type LogFile struct {
	Path string
	Info os.FileInfo
}

func (lu *LogUnit) getLogFiles() ([]*LogFile, error) {

	pattern := lu.FilePattern
	if !filepath.IsAbs(pattern) {
		if wd, err := os.Getwd(); nil == err {
			pattern = wd + string(os.PathSeparator) + pattern
		} else {
			return nil, err
		}
	}

	paths, err := filepath.Glob(pattern)
	if nil != err {
		return nil, err
	}

	result := make([]*LogFile, 0, len(paths))
	for _, path := range paths {

		stat, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		li := &LogFile{
			Path: path,
			Info: stat,
		}

		if !stat.Mode().IsDir() {
			result = append(result, li)
		}
	}

	sort.Slice(result[:], func(i, j int) bool {
		return result[i].getMTime() > result[j].getMTime()
	})

	return result, nil
}

func (li *LogFile) getMTime() int64 {
	return li.Info.ModTime().Unix()
}
