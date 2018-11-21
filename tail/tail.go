package tail

import (
	"os"
	"strings"
	"context"
	"github.com/iobestar/logship/utils/logger"
)

func ReadTail(ctx context.Context, file *os.File) (<-chan string, error) {
	fileInfo, err := file.Stat()
	if nil != err {
		return nil, err
	}

	fileSize := fileInfo.Size()
	bufSize := int64(4096)
	if fileSize < bufSize {
		bufSize = fileSize
	}

	var (
		buf    = make([]byte, bufSize)
		total  = int64(0)
		remain = fileSize
		index  = int64(0)
		lines = make(chan string, 32)
	)

	go func() {
		defer close(lines)
		var line = strings.Builder{}
		for {
			if remain == 0 {
				lines <- reverse(line.String())
				break
			}

			if int64(cap(buf)) >= remain {
				index = 0
			} else {
				index = fileInfo.Size() - total - int64(cap(buf))
			}

			n, err := file.ReadAt(buf, index)
			if nil != err {
				logger.Error.Printf("Error reading file: %s", err.Error())
				return
			}

			if int64(cap(buf)) >= remain {
				n = int(remain)
			}

			for i := n; i > 0; i = i - 1 {
				c := buf[i-1]
				if c == '\n' {
					select {
					case lines <- reverse(line.String()):
						line.Reset()
					case <-ctx.Done():
						return
					}
				} else {
					line.WriteByte(c)
				}
			}

			total += int64(n)
			remain -= int64(n)
		}
	}()

	return lines, nil
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
