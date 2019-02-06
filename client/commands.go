package client

import (
	"context"
	"fmt"
	"github.com/iobestar/logship/config"
	"github.com/iobestar/logship/unit/rpc"
	"github.com/iobestar/logship/utils/logger"
	"google.golang.org/grpc"
	"io"
)

func Lines(ctx context.Context, conn *grpc.ClientConn, unitId string, count int) error {
	service := rpc.NewLogUnitServiceClient(conn)
	lineStream, err := service.NLines(ctx, &rpc.NLineRQ{
		UnitId: unitId,
		Count:  int32(count),
	})
	if nil != err {
		return err
	}

	var result []string
	for {
		rs, err := lineStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error.Fatal(err)
		}

		result = append(result, rs.Line)
	}
	for i := len(result); i > 0; i = i - 1 {
		fmt.Println(result[i-1])
	}
	return nil
}

func NLogs(ctx context.Context, conn *grpc.ClientConn, unitId string, count int, logReaderConfig config.LogReaderConfig) error {
	service := rpc.NewLogUnitServiceClient(conn)
	logStream, err := service.NLogs(ctx, &rpc.NLogRQ{
		UnitId:         unitId,
		Count:          int32(count),
		DateTimeLayout: logReaderConfig.DateTimeLayout,
		LogPattern:     logReaderConfig.LogPattern,
	})
	if nil != err {
		return err
	}

	var result []string
	for {
		rs, err := logStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error.Fatal(err)
		}

		result = append(result, rs.Log)
	}
	for i := len(result); i > 0; i = i - 1 {
		fmt.Println(result[i-1])
	}
	return nil
}

func TLogs(ctx context.Context, conn *grpc.ClientConn, unitId string, duration string, logReaderConfig config.LogReaderConfig) error {
	service := rpc.NewLogUnitServiceClient(conn)
	logStream, err := service.TLogs(ctx, &rpc.TLogRQ{
		UnitId:         unitId,
		Duration:       duration,
		DateTimeLayout: logReaderConfig.DateTimeLayout,
		LogPattern:     logReaderConfig.LogPattern,
	})
	if nil != err {
		return err
	}

	var result []string
	for {
		rs, err := logStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error.Fatal(err)
		}

		result = append(result, rs.Log)
	}
	for i := len(result); i > 0; i = i - 1 {
		fmt.Println(result[i-1])
	}
	return nil
}
