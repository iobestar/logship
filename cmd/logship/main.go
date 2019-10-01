package main

import (
	"context"
	cl "github.com/iobestar/logship/client"
	"github.com/iobestar/logship/config"
	"github.com/iobestar/logship/unit"
	pb "github.com/iobestar/logship/unit/rpc"
	"github.com/iobestar/logship/utils/logger"
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

var (
	app = kingpin.New("logship", "Tool for shipping logs")

	// server
	server   = app.Command("server", "Logship server mode")
	address  = server.Flag("address", "Logship server address").Default(":" + strconv.Itoa(defaultPort)).String()
	logUnits = server.Flag("logunits", "Logship server log units").Envar("LOG_UNITS").String()

	// client
	client       = app.Command("client", "Logship client mode").Default()
	configPath   = client.Flag("config", "Configuration path").Default("logship.yml").String()
	target       = client.Flag("target", "Logship server address").Short('t').Default(defaultTarget).String()
	units        = client.Command("units", "List log units").Default()
	nLog         = client.Command("nlogs", "Fetch fixed number of logs from Logship server")
	nLogUnitId   = nLog.Arg("unit-id", "Log unit identifier").Required().String()
	nLogCount    = nLog.Arg("count", "Number of log records").Required().Int()
	nLogReaderId = nLog.Arg("reader-id", "Log reader identifier").Default("").String()
	tLog         = client.Command("tlogs", "Fetch logs by time from Logship server")
	tLogUnitId   = tLog.Arg("unit-id", "Log unit identifier").Required().String()
	tLogDuration = tLog.Arg("duration", "Time duration (e.g. 1h, 45m").Required().String()
	tLogReaderId = tLog.Arg("reader-id", "Log reader identifier").Default("").String()
	nLine        = client.Command("nlines", "Fetch fixed numbers of line form Logship server")
	nLineUnitId  = nLine.Arg("unit-id", "Log unit identifier").Required().String()
	nLineCount   = nLine.Arg("count", "Number of lines").Int()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case server.FullCommand():

		logger.Info.Println("Logship starting in mode: SERVER")

		logUnits, err := unit.NewLogUnits(*logUnits)
		if nil != err {
			logger.Error.Fatal(err)
		}
		logger.Info.Printf("Log units: %v", logUnits.GetLogUnitIds())

		logService, err := pb.NewLogService(logUnits)
		if nil != err {
			logger.Error.Fatal(err)
		}

		grpcS := grpc.NewServer()

		pb.RegisterLogUnitServiceServer(grpcS, logService)

		lis, err := net.Listen("tcp", *address)
		if err != nil {
			logger.Error.Fatal(err)
		}

		m := cmux.New(lis)
		grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		httpL := m.Match(cmux.HTTP1Fast())

		go func() {
			logger.Info.Printf("Starting RPC server: %s", *address)
			if err = grpcS.Serve(grpcL); err != cmux.ErrListenerClosed && err != nil {
				logger.Error.Fatal(err)
			}
		}()

		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		httpS := &http.Server{
			Handler: mux,
		}

		go func() {
			logger.Info.Printf("Starting HTTP server: %s", *address)
			if err = httpS.Serve(httpL); err != http.ErrServerClosed {
				logger.Error.Fatal(err)
			}
		}()

		go func() {
			if err := m.Serve(); !strings.Contains(err.Error(), "use of closed network connection") {
				panic(err)
			}
		}()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sig:
			err = httpS.Shutdown(context.Background())
			if nil != err {
				logger.Warning.Println(err)
			}
			grpcS.GracefulStop()
		}
	case units.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				logger.Warning.Println(err)
			}
		}()

		err := cl.Units(context.Background(), conn)
		if nil != err {
			logger.Error.Fatalf("error executing units command: %s", err.Error())
		}
	case nLine.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				logger.Warning.Println(err)
			}
		}()

		err := cl.NLines(context.Background(), conn, *nLineUnitId, *nLineCount)
		if nil != err {
			logger.Error.Fatalf("error executing units command: %s", err.Error())
		}
	case nLog.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				logger.Warning.Println(err)
			}
		}()

		cfg, err := config.ParseConfig(*configPath)
		if nil != err {
			logger.Error.Fatalf("error parsing configuration: %s", err.Error())
		}

		logReaderCfg := cfg.GetLogReaderConfig(*nLogReaderId)
		err = cl.NLogs(context.Background(), conn, *nLogUnitId, *nLogCount, logReaderCfg)
		if nil != err {
			logger.Error.Fatalf("error executing nlogs command: %s", err.Error())
		}
	case tLog.FullCommand():
		conn := connect(adjustTarget(*target))
		defer func() {
			if err := conn.Close(); nil != err {
				logger.Warning.Println(err)
			}
		}()

		cfg, err := config.ParseConfig(*configPath)
		if nil != err {
			logger.Error.Fatalf("error parsing configuration: %s", err.Error())
		}

		logReaderCfg := cfg.GetLogReaderConfig(*tLogReaderId)
		err = cl.TLogs(context.Background(), conn, *tLogUnitId, *tLogDuration, logReaderCfg)
		if nil != err {
			logger.Error.Fatalf("error executing nlogs command: %s", err.Error())
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
		logger.Error.Fatalf("target is empty")
	}
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		logger.Error.Fatalf("unable to create channel for target %s: %s", target, err.Error())
	}
	return conn
}
