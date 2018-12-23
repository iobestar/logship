package unit

import (
	"bytes"
	"context"
	"regexp"
	"strings"
	"time"
)

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

func (lu *LogUnit) createLog(params map[string]string, payload []string, dateTimeLayout string) (*LogEntry, error) {

	var dateTime string
	var timestamp int64
	var err error
	if value, valueExist := params["datetime"]; valueExist {
		dateTime = value
		timestamp, err = toTimestamp(dateTimeLayout, dateTime)
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

func (lu *LogUnit) getLogParameters(match []string, logPattern *regexp.Regexp) map[string]string {

	result := make(map[string]string)
	for i, name := range logPattern.SubexpNames() {
		if i > 0 && i <= len(match) {
			result[name] = match[i]
		}
	}
	return result
}

func (lu *LogUnit) StreamLogs(ctx context.Context, dateTimeLayout string, logPattern *regexp.Regexp) (chan *LogEntry, chan error) {

	streamCtx, _ := context.WithCancel(ctx)
	lines, linesErrors := lu.StreamLines(streamCtx)

	logs := make(chan *LogEntry, 32)
	errors := make(chan error, 1)
	go func() {
		defer func() {
			close(logs)
			close(errors)
		}()
		var logPayload []string
		for {
			select {
			case line, ok := <-lines:
				if !ok {
					return
				}
				logPayload = append(logPayload, line)
				match := logPattern.FindStringSubmatch(line)
				if len(match) > 0 && len(logPayload) > 0 {
					log, err := lu.createLog(lu.getLogParameters(match, logPattern), logPayload, dateTimeLayout)
					if nil != err {
						errors <- err
						return
					}
					logs <- log
					logPayload = make([]string, 0)
				}
			case err, ok := <-linesErrors:
				if ok {
					errors <- err
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return logs, errors
}
