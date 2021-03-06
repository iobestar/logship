package tail

import (
	"context"
	"os"
	"strings"
)

const lineBuffer = 256

func ReadTail(ctx context.Context, file *os.File) (<-chan string, <- chan error) {

	var (
		lines  = make(chan string, lineBuffer)
		errors  = make(chan error)
	)

	fileInfo, err := file.Stat()
	if nil != err {
		errors <- err
		close(lines)
		close(errors)
		return lines, errors
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
	)

	go func() {
		defer close(lines)
		defer close(errors)
		var lineBuffer = strings.Builder{}
		for {
			if remain == 0 {
				lines <- reverse(lineBuffer.String())
				return
			}

			if int64(cap(buf)) >= remain {
				index = 0
			} else {
				index = fileInfo.Size() - total - int64(cap(buf))
			}

			n, err := file.ReadAt(buf, index)
			if nil != err {
				errors <- err
				return
			}

			if int64(cap(buf)) >= remain {
				n = int(remain)
			}

			for i := n; i > 0; i = i - 1 {
				c := buf[i-1]
				if c == '\n' {
					if i == n && total == 0 { // is last new line
						continue
					}
					select {
					case lines <- reverse(lineBuffer.String()):
						lineBuffer.Reset()
					case <-ctx.Done():
						return
					}
				} else {
					lineBuffer.WriteByte(c)
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
