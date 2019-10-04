package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log/level"
	cl "github.com/iobestar/logship/client"
	"github.com/iobestar/logship/config"
	"github.com/iobestar/logship/logging"
	"github.com/iobestar/logship/unit"
	pb "github.com/iobestar/logship/unit/rpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

var (
	defaultPort   = 11034
	defaultTarget = "localhost:" + strconv.Itoa(defaultPort)
)

const (
	defaultLogPattern     = "^(?P<datetime>\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}\\.\\d{3}).*"
	defaultDateTimeLayout = "2006-01-02 15:04:05.000"
)

var (
	app = kingpin.New("logship", "Tool for shipping logs")

	// version
	versionCmd = app.Command("version", "")

	// server
	server     = app.Command("server", "Logship server mode")
	address    = server.Flag("address", "Logship server address").Default(":" + strconv.Itoa(defaultPort)).String()
	configFile = server.Flag("config", "Configuration file").Default("logship.yml").String()

	// client
	client         = app.Command("client", "Logship client mode").Default()
	logPattern     = client.Flag("log-pattern", "Regex pattern of log in log files").Short('p').Envar("LOGSHIP_LOG_PATTERN").Default(defaultLogPattern).String()
	dateTimeLayout = client.Flag("datetime-layout", "Date time layout of log in log files (https://golang.org/pkg/time/#Time.Format)").Short('l').Envar("LOGSHIP_DATETIME_LAYOUT").Default(defaultDateTimeLayout).String()
	target         = client.Flag("target", "Logship server address").Short('t').Default(defaultTarget).String()
	units          = client.Command("units", "List log units").Default()
	nLog           = client.Command("nlogs", "Fetch fixed number of logs from Logship server")
	nLogUnitId     = nLog.Arg("unit-id", "Log unit identifier").Required().String()
	nLogCount      = nLog.Arg("count", "Number of log records").Required().Int()
	tLog           = client.Command("tlogs", "Fetch logs by time from Logship server")
	tLogUnitId     = tLog.Arg("unit-id", "Log unit identifier").Required().String()
	tLogDuration   = tLog.Arg("duration", "Time duration (e.g. 1h, 45m").Required().String()
	nLine          = client.Command("nlines", "Fetch fixed numbers of line from Logship server")
	nLineUnitId    = nLine.Arg("unit-id", "Log unit identifier").Required().String()
	nLineCount     = nLine.Arg("count", "Number of lines").Int()
)

var (
	Version  string
	Revision string
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case versionCmd.FullCommand():
		fmt.Printf("%-8s: %s\n", "version", Version)
		fmt.Printf("%s: %s\n", "revision", Revision)
	case server.FullCommand():
		cfg, err := config.ParseConfig(*configFile)
		if err != nil {
			logging.StdErrLogger.Fatal(err)
		}

		level.Info(logging.Logger).Log("msg", "logship starting in server mode")

		logUnits, err := unit.NewLogUnits(cfg.LogUnits)
		if nil != err {
			logging.StdErrLogger.Fatal(err)
		}
		level.Info(logging.Logger).Log("log_units", fmt.Sprintf("%v", logUnits.GetLogUnitIds()))

		logService, err := pb.NewLogService(logUnits)
		if nil != err {
			logging.StdErrLogger.Fatal(err)
		}

		grpcS := grpc.NewServer()

		pb.RegisterLogUnitServiceServer(grpcS, logService)

		lis, err := net.Listen("tcp", *address)
		if err != nil {
			logging.StdErrLogger.Fatal(err)
		}

		m := cmux.New(lis)
		grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		httpL := m.Match(cmux.HTTP1Fast())

		go func() {
			level.Info(logging.Logger).Log("msg", "starting RPC server", "address", *address)
			if err = grpcS.Serve(grpcL); err != cmux.ErrListenerClosed && err != nil {
				logging.StdErrLogger.Fatal(err)
			}
		}()

		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		httpS := &http.Server{
			Handler: mux,
		}

		go func() {
			level.Info(logging.Logger).Log("msg", "starting HTTP server", "address", *address)
			if err = httpS.Serve(httpL); err != http.ErrServerClosed && err != nil {
				logging.StdErrLogger.Fatal(err)
			}
		}()

		go func() {
			if err := m.Serve(); err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
				logging.StdErrLogger.Fatal(err)
			}
		}()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sig:
			err = httpS.Shutdown(context.Background())
			if nil != err {
				level.Warn(logging.Logger).Log("err", err.Error())
			}
			grpcS.GracefulStop()
		}
	case units.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				level.Warn(logging.Logger).Log("err", err.Error())
			}
		}()

		err := cl.Units(context.Background(), conn)
		if nil != err {
			level.Error(logging.Logger).Log("msg", "error executing command", "command", "units", "err", err.Error())
		}
	case nLine.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				level.Warn(logging.Logger).Log("err", err.Error())
			}
		}()

		err := cl.NLines(context.Background(), conn, *nLineUnitId, *nLineCount)
		if nil != err {
			level.Error(logging.Logger).Log("msg", "error executing command", "command", "nlines", "err", err.Error())
		}
	case nLog.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				level.Warn(logging.Logger).Log("err", err.Error())
			}
		}()

		err := cl.NLogs(context.Background(), conn, *nLogUnitId, *nLogCount, *dateTimeLayout, *logPattern)
		if nil != err {
			level.Error(logging.Logger).Log("msg", "error executing command", "command", "nlogs", "err", err.Error())
		}
	case tLog.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				level.Warn(logging.Logger).Log("err", err.Error())
			}
		}()

		err := cl.TLogs(context.Background(), conn, *tLogUnitId, *tLogDuration, *dateTimeLayout, *logPattern)
		if nil != err {
			level.Error(logging.Logger).Log("msg", "error executing command", "command", "tlogs", "err", err.Error())
		}
	default:
	}
}

func adjustTarget(target string) string {

	if target == defaultTarget {
		return target
	}

	if len(target) == 0 {
		return defaultTarget
	}

	if strings.Index(target, ":") == -1 {
		return target + ":" + strconv.Itoa(defaultPort)
	}
	return target
}

func connect(target string) *grpc.ClientConn {
	if len(target) == 0 {
		logging.StdErrLogger.Fatal("target is empty")

	}
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		logging.StdErrLogger.Fatalf("unable to create channel for target %s: %s", target, err.Error())
	}
	return conn
}
