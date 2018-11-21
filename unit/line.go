package unit

import (
	"os"
	"github.com/iobestar/logship/tail"
	"context"
)

func (lu *LogUnit) getLines(consume func(count int, line string) (bool, error)) error {

	logFiles, err := lu.getLogFiles()
	if nil != err {
		return err
	}

	if len(logFiles) == 0 {
		return nil
	}

	files := make([]*os.File, 0, len(logFiles))
	defer func() {
		for _, f := range files {
			f.Close()
		}
	}()

	for _, logFile := range logFiles {

		logPath := logFile.Path
		if len(logPath) == 0 {
			return nil
		}

		file, err := os.Open(logPath)
		if nil != err {
			return err
		}
		files = append(files, file)

		ctx, cancel := context.WithCancel(context.Background())
		lines, err := tail.ReadTail(ctx, file)
		if nil != err {
			return nil
		}

		count := 0
		for {
			logLine, more := <-lines
			if more {
				count = count + 1
				done, err := consume(count, logLine)

				if nil != err {
					return err
				}

				if done {
					break
				}
			} else {
				break
			}
		}
		cancel()
	}

	return nil
}
