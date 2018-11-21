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
	fLine        = client.Command("flines", "Stream lines from Logship server")
	fLineUnitId  = fLine.Arg("unit-id", "Log unit identifier").Required().String()
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

		logger.Info.Printf("Starting RPC server: %s", *address)
		if err = grpcServer.Serve(lis); err != nil {
			logger.Error.Fatal(err)
		}
	case client.FullCommand():
		fmt.Println("CLIENT")
	case units.FullCommand():
		cli, err := NewLogshipClient((*targets)[0])
		if nil != err {
			panic(err)
		}
		defer cli.Close()

		unitStream, err := cli.GetUnits(context.Background(), &pb.Empty{})
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

		cli, err := NewLogshipClient((*targets)[0])
		if nil != err {
			panic(err)
		}
		defer cli.Close()

		logStream, err := cli.GetNLogs(context.Background(), &pb.NLogRQ{
			UnitId: *nLogUnitId,
			Count:  int32(*nLogCount),
		})
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

		cli, err := NewLogshipClient((*targets)[0])
		if nil != err {
			panic(err)
		}
		defer cli.Close()

		logStream, err := cli.GetTLogs(context.Background(), &pb.TLogRQ{
			UnitId:   *tLogUnitId,
			Duration: *tLogDuration,
			Offset:   0,
		})

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
		cli, err := NewLogshipClient((*targets)[0])
		if nil != err {
			panic(err)
		}
		defer cli.Close()

		rq := &pb.NLineRQ{
			UnitId: *nLineUnitId,
			Count:  int32(*nLineCount),
		}

		lineStream, err := cli.GetNLines(context.Background(), rq)
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
	case fLine.FullCommand():
		cli, err := NewLogshipClient((*targets)[0])
		if nil != err {
			panic(err)
		}
		defer cli.Close()

		rq := &pb.FLineRQ{
			UnitId: *fLineUnitId,
		}
		ctx, cancel := context.WithCancel(context.Background())
		lineStream, err := cli.GetFLines(ctx, rq)
		go func() {
			for {
				rs, err := lineStream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					panic(err)
				}
				fmt.Print(rs.Line)
			}
		}()

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sig:
			cancel()
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

func (lc *LogshipClient) GetFLines(ctx context.Context, rq *pb.FLineRQ, opts ...grpc.CallOption) (pb.LogUnitService_GetNLinesClient, error) {
	return lc.delegate.GetFLines(ctx, rq, opts...)
}

func (lc *LogshipClient) Close() error {
	return lc.cnn.Close()
}

func NewLogshipClient(target string) (*LogshipClient, error) {

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	service := pb.NewLogUnitServiceClient(conn)
	return &LogshipClient{
		cnn:      conn,
		delegate: service,
	}, nil
}
