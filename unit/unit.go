package unit

import (
	"context"
	"github.com/iobestar/logship/tail"
	"github.com/iobestar/logship/utils/logger"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type LogUnit struct {
	Id             string
	FilePattern    string
}

type LogFile struct {
	Path string
	Info os.FileInfo
}

type LogUnits map[string]*LogUnit

func NewLogUnits(units []string) LogUnits {

	result := LogUnits{}
	for _, _unit := range units {
		_unit := strings.Split(_unit, ":")
		unit := &LogUnit{
			Id: _unit[0],
			FilePattern: _unit[1],
		}
		result[unit.Id] = unit
	}
	return result
}

func (lu LogUnits) GetLogUnit(unitId string) *LogUnit {

	if u, ok := lu[unitId]; ok {
		return u
	}
	return nil
}

func (lu LogUnits) GetLogUnitIds() []string {

	var result []string
	for id := range lu {
		result = append(result, id)
	}
	return result
}

func (lu *LogUnit) StreamLines(ctx context.Context) (<-chan string, <-chan error) {

	lines := make(chan string, 256)
	errors := make(chan error, 1)
	go func() {

		defer func() {
			close(errors)
			close(lines)
		}()

		logFiles, err := lu.getLogFiles()
		if nil != err {
			errors <- err
			return
		}

		if len(logFiles) == 0 {
			return
		}

		files := make([]*os.File, 0, len(logFiles))
		defer func() {
			for _, f := range files {
				err := f.Close()
				if nil != err {
					logger.Warning.Println(err)
				}
			}
		}()

		for _, logFile := range logFiles {

			logPath := logFile.Path
			if len(logPath) == 0 {
				return
			}

			file, err := os.Open(logPath)
			if nil != err {
				errors <- err
				return
			}
			files = append(files, file)

			tailLines, err := tail.ReadTail(ctx, file)
			if nil != err {
				errors <- err
				return
			}
		file:
			for {
				select {
				case line, ok := <-tailLines:
					if ok {
						lines <- line
					} else {
						break file
					}
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return lines, errors
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
