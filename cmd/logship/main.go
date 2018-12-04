package main

import (
	"os"
	pb "github.com/iobestar/logship/unit/rpc"
	"gopkg.in/alecthomas/kingpin.v2"
	"google.golang.org/grpc"
	"net"
	"fmt"
	"context"
	"io"
	"github.com/iobestar/logship/unit"
	"github.com/iobestar/logship/config"
	"syscall"
	"os/signal"
	"github.com/iobestar/logship/utils/logger"
	"github.com/soheilhy/cmux"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"strings"
)

var (
	app = kingpin.New("logship", "Tool for shipping logs")

	configPath = app.Flag("config", "Configuration path").Default("logship.yml").String()

	// server
	server  = app.Command("server", "Logship server mode")
	address = server.Flag("address", "Logship server address").Default("0.0.0.0:3340").String()

	// client
	client       = app.Command("client", "Logship client mode").Default()
	targets      = client.Flag("target", "Logship server addresses").Default("localhost:3340").Strings()
	units        = client.Command("units", "List log units").Default()
	nLog         = client.Command("nlogs", "Fetch fixed number of logs from Logship server")
	nLogUnitId   = nLog.Arg("unit-id", "Log unit identifier").Required().String()
	nLogCount    = nLog.Arg("count", "Number of log records").Required().Int()
	tLog         = client.Command("tlogs", "Fetch logs by time from Logship server")
	tLogUnitId   = tLog.Arg("unit-id", "Log unit identifier").Required().String()
	tLogDuration = tLog.Arg("duration", "Time duration (e.g. 1h, 45m").Required().String()
	nLine        = client.Command("nlines", "Fetch fixed numbers of line form Logship server")
	nLineUnitId  = nLine.Arg("unit-id", "Log unit identifier").Required().String()
	nLineCount   = nLine.Arg("count", "Number of lines").Int()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case server.FullCommand():

		logger.Info.Println("Logship starting in mode: SERVER")

		cfg, err := config.Parse(*configPath)
		if nil != err {
			logger.Error.Fatal(err)
		}
		logger.Info.Printf("Parsed configuration: %s", *configPath)

		unitManager, err := unit.NewManager(*cfg)
		if nil != err {
			logger.Error.Fatal(err)
		}
		logger.Info.Printf("Created log unit manager: %v", unitManager.GetLogUnitIds())

		logService, err := pb.NewLogService(unitManager)
		if nil != err {
			logger.Error.Fatal(err)
		}

		grpcServer := grpc.NewServer()

		pb.RegisterLogUnitServiceServer(grpcServer, logService)

		lis, err := net.Listen("tcp", *address)
		if err != nil {
			logger.Error.Fatal(err)
		}

		m := cmux.New(lis)
		grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
		httpL := m.Match(cmux.HTTP1Fast())

		go func() {
			logger.Info.Printf("Starting RPC server: %s", *address)
			if err = grpcServer.Serve(grpcL); err != cmux.ErrListenerClosed && err != nil {
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
			httpS.Shutdown(context.Background())
			grpcServer.GracefulStop()
		}
	case units.FullCommand():
		cli := NewLogshipClient((*targets)[0])
		defer cli.Close()

		unitStream, err := cli.GetUnits(context.Background(), &pb.Empty{})
		if nil!= err {
			logger.Error.Fatalf("Error executing units command: %s", err.Error())
		}
		for {
			u, err := unitStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(u.Unit)
		}
	case nLog.FullCommand():

		cli := NewLogshipClient((*targets)[0])
		defer cli.Close()

		logStream, err := cli.GetNLogs(context.Background(), &pb.NLogRQ{
			UnitId: *nLogUnitId,
			Count:  int32(*nLogCount),
		})
		if nil!= err {
			logger.Error.Fatalf("Error executing nlogs command: %s", err.Error())
		}

		var result []string
		for {
			logEntry, err := logStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			result = append(result, logEntry.Payload)
		}
		for i := len(result); i > 0; i = i - 1 {
			fmt.Println(result[i-1])
		}
	case tLog.FullCommand():

		cli := NewLogshipClient((*targets)[0])
		defer cli.Close()

		logStream, err := cli.GetTLogs(context.Background(), &pb.TLogRQ{
			UnitId:   *tLogUnitId,
			Duration: *tLogDuration,
			Offset:   0,
		})
		if nil!= err {
			logger.Error.Fatalf("Error executing tlogs command: %s", err.Error())
		}

		var result []string
		for {
			logEntry, err := logStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			result = append(result, logEntry.Payload)
		}
		for i := len(result); i > 0; i = i - 1 {
			fmt.Println(result[i-1])
		}
	case nLine.FullCommand():
		cli := NewLogshipClient((*targets)[0])
		defer cli.Close()

		lineStream, err := cli.GetNLines(context.Background(), &pb.NLineRQ{
			UnitId: *nLineUnitId,
			Count:  int32(*nLineCount),
		})
		if nil!= err {
			logger.Error.Fatalf("Error executing nlines command: %s", err.Error())
		}

		var result []string
		for {
			rs, err := lineStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			result = append(result, rs.Line)
		}
		for i := len(result); i > 0; i = i - 1 {
			fmt.Println(result[i-1])
		}
	default:
	}
}

type LogshipClient struct {
	cnn      *grpc.ClientConn
	delegate pb.LogUnitServiceClient
}

func (lc *LogshipClient) GetNLogs(ctx context.Context, in *pb.NLogRQ, opts ...grpc.CallOption) (pb.LogUnitService_GetNLogsClient, error) {
	return lc.delegate.GetNLogs(ctx, in, opts...)
}

func (lc *LogshipClient) GetTLogs(ctx context.Context, in *pb.TLogRQ, opts ...grpc.CallOption) (pb.LogUnitService_GetTLogsClient, error) {
	return lc.delegate.GetTLogs(ctx, in, opts...)
}

func (lc *LogshipClient) GetUnits(ctx context.Context, rq *pb.Empty, opts ...grpc.CallOption) (pb.LogUnitService_GetUnitsClient, error) {
	return lc.delegate.GetUnits(ctx, rq, opts...)
}

func (lc *LogshipClient) GetNLines(ctx context.Context, rq *pb.NLineRQ, opts ...grpc.CallOption) (pb.LogUnitService_GetNLinesClient, error) {
	return lc.delegate.GetNLines(ctx, rq, opts...)
}

func (lc *LogshipClient) Close() error {
	return lc.cnn.Close()
}

func NewLogshipClient(target string) (*LogshipClient) {

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		logger.Error.Fatalf("Unable to create logship client: %s", err.Error())
	}

	service := pb.NewLogUnitServiceClient(conn)
	return &LogshipClient{
		cnn:      conn,
		delegate: service,
	}
}
