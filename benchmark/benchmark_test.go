package benchmark

import (
	"testing"
	"github.com/iobestar/logship/tail"
	"os"
	"context"
)

// SH: go test -bench=. -benchmem -cpuprofile=cpu.out -memprofile=mem.out benchmark_test.go
func BenchmarkTailing(b *testing.B) {

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f, _ := os.Open("fixture/output.log")
		lines, _ := tail.ReadTail(context.Background(), f)
		b.StartTimer()
		for {
			_, more := <-lines
			if more {
			} else {
				break
			}
		}
		b.StopTimer()
		f.Close()
	}
}
