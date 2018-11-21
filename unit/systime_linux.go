package unit

import (
	"syscall"
	"time"
)

func (li *LogFile) GetATime() int64 {

	stat := li.Info.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat.Atim.Sec), int64(stat.Atim.Nsec)).Unix()
}

func (li *LogFile) GetCTime() int64 {

	stat := li.Info.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec)).Unix()
}
