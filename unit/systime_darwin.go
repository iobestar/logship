package unit

import (
	"syscall"
	"time"
)

func (li *LogFile) GetATime() int64 {

	stat := li.Info.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat.Atimespec.Sec), int64(stat.Atimespec.Nsec)).Unix()
}

func (li *LogFile) GetCTime() int64 {

	stat := li.Info.Sys().(*syscall.Stat_t)
	return time.Unix(int64(stat.Ctimespec.Sec), int64(stat.Ctimespec.Nsec)).Unix()
}
