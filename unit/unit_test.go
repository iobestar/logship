package unit

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
	"time"
)

func TestLogUnit_StreamLines(t *testing.T) {

	unit := &LogUnit{
		Id:          "test",
		FilePattern: "fixture/lines.log",
	}

	result := make([]string, 0, 0)
	lines, errors := unit.StreamLines(context.Background())

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
				assert.Nil(t, err, "Unexpected error during lines streaming")
			}
		}
	}

	assert.EqualValues(t, 4, len(result))
	assert.EqualValues(t, "third line", result[0])
	assert.EqualValues(t, "", result[1])
	assert.EqualValues(t, "second line", result[2])
}

func TestLogUnit_StreamLines_MultipleFiles(t *testing.T) {
	unit := &LogUnit{
		Id:          "test",
		FilePattern: "fixture/lines_multiple.log*",
	}

	result := make([]string, 0, 0)
	lines, errors := unit.StreamLines(context.Background())

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
				assert.Nil(t, err, "Unexpected error during lines streaming")
			}
		}
	}

	assert.EqualValues(t, 11, len(result))
}

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