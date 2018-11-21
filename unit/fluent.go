package unit

import (
	"path"
	"github.com/fsnotify/fsnotify"
	"context"
	"os"
	"github.com/iobestar/logship/utils/logger"
	"strings"
	"io"
)

func (lu *LogUnit) FluentRead(ctx context.Context) (<-chan string, error) {

	result := make(chan string)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	dir, _ := path.Split(lu.FilePattern)
	err = watcher.Add(dir)
	if nil != err {
		return nil, err
	}

	go func() {
		defer watcher.Close()
		defer close(result)

		var (
			file, index = lu.getLatestFileWithSize()
			buf         = make([]byte, 4096)
			output      = strings.Builder{}
		)

		defer func() {
			if nil != file {
				file.Close()
			}
		}()
		for {
			select {
			case <-ctx.Done():
				watcher.Close()
				logger.Info.Printf("Fluent read closed for: %s", lu.FilePattern)
				return
			case event, more := <-watcher.Events:
				if !more {
					break
				}

				if event.Op&fsnotify.Write == fsnotify.Write {

					if nil == file && index == -1 {
						file, index = lu.getLatestFileWithSize()
					}

					if nil != file && strings.HasSuffix(file.Name(), event.Name) {
						read:
						for {
							if index == -1 {
								break
							}

							n, err := file.ReadAt(buf, index)
							if n > 0 {
								output.Write(buf[:n])
								select {
								case <-ctx.Done():
									break read
								case result <- output.String():
									output.Reset()
								}
								index = index + int64(n)
							}

							if nil != err && err == io.EOF {
								break
							} else {
								logger.Error.Println(err)
								return
							}

							select {
							case <-ctx.Done():
								break read
							default:
							}
						}
					}
				}

				if event.Op&fsnotify.Rename == fsnotify.Rename {
					if nil != file && strings.HasSuffix(file.Name(), event.Name) {
						file.Close()
						file, index = nil, -1
					}
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					if nil != file && strings.HasSuffix(file.Name(), event.Name) {
						file.Close()
						file, index = nil, -1
					}
				}
			case err, ok := <-watcher.Errors:
				if nil != err {
					logger.Error.Println(err)
				}
				if !ok {
					return
				}
			}
		}
	}()
	return result, nil
}

func (lu *LogUnit) getLatestFileWithSize() (*os.File, int64) {

	logFiles, err := lu.getLogFiles()
	if nil != err {
		logger.Error.Println(err)
		return nil, -1
	}

	if len(logFiles) == 0 {
		return nil, -1
	}

	file, err := os.Open(logFiles[0].Path)
	if nil != err {
		logger.Error.Println(err)
		return nil, -1
	}

	stat, err := file.Stat()
	if nil != err {
		logger.Error.Println(err)
		return nil, -1
	}

	return file, stat.Size()
}
