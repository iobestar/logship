package unit

import (
	"testing"
	"os"
	"context"
	"time"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"io/ioutil"
	"path"
	"strings"
)

func TestFluent(t *testing.T) {

	dir, _ := ioutil.TempDir("", "logshiptest")
	defer os.RemoveAll(dir)
	file, _ := os.Create(path.Join(dir, "test.log"))

	ctx, _ := context.WithTimeout(context.Background(), 1 * time.Second)

	f, _ := os.OpenFile(file.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	tick := time.NewTicker(10 * time.Millisecond)
	i := 0;
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			case <-tick.C:
				i = i + 1
				_, err := f.WriteString("line" + strconv.Itoa(i))
				if nil != err {
					fmt.Println(err.Error())
				}
			}
		}
	}()

	unit := &LogUnit{
		Id:          "test",
		FilePattern: path.Join(dir, "test*.log"),
	}

	data, err := unit.FluentRead(ctx)
	assert.Nil(t, err, "Fluent reading error")

	result := strings.Builder{}
	for {
		d, more := <-data
		if more {
			result.WriteString(d)
		} else {
			break
		}
	}

	assert.True(t, strings.Contains(result.String(), "line10"))
	assert.True(t, strings.Contains(result.String(), "line20"))
	assert.True(t, strings.Contains(result.String(), "line30"))
	assert.True(t, strings.Contains(result.String(), "line40"))
	assert.True(t, strings.Contains(result.String(), "line50"))
	assert.True(t, strings.Contains(result.String(), "line60"))
	assert.True(t, strings.Contains(result.String(), "line70"))
	assert.True(t, strings.Contains(result.String(), "line90"))
	assert.True(t, strings.Contains(result.String(), "line90"))
}
