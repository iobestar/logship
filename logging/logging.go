package logging

import (
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	stdlog "log"
	"os"
)

var (
	Logger       kitlog.Logger
	StdErrLogger *stdlog.Logger
)

func init() {
	Logger = kitlog.NewLogfmtLogger(kitlog.NewSyncWriter(os.Stdout))
	Logger = kitlog.With(Logger, "ts", kitlog.DefaultTimestampUTC, "caller", kitlog.DefaultCaller)
	StdErrLogger = stdlog.New(kitlog.NewStdlibAdapter(level.Error(Logger)), "", 0)
}
