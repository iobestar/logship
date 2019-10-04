package client

import (
	"context"
	"fmt"
	"github.com/iobestar/logship/logging"
	"github.com/iobestar/logship/unit/rpc"
	"google.golang.org/grpc"
	"io"
)

func NLines(ctx context.Context, conn *grpc.ClientConn, unitId string, count int) error {
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
			logging.StdErrLogger.Fatal(err)
		}

		result = append(result, rs.Line)
	}
	for i := len(result); i > 0; i = i - 1 {
		fmt.Println(result[i-1])
	}
	return nil
}

func NLogs(ctx context.Context, conn *grpc.ClientConn, unitId string, count int, dateTimeLayout, logPattern string) error {
	service := rpc.NewLogUnitServiceClient(conn)
	logStream, err := service.NLogs(ctx, &rpc.NLogRQ{
		UnitId:         unitId,
		Count:          int32(count),
		DateTimeLayout: dateTimeLayout,
		LogPattern:     logPattern,
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
			logging.StdErrLogger.Fatal(err)
		}

		result = append(result, rs.Log)
	}
	for i := len(result); i > 0; i = i - 1 {
		fmt.Println(result[i-1])
	}
	return nil
}

func TLogs(ctx context.Context, conn *grpc.ClientConn, unitId string, duration string, dateTimeLayout, logPattern string) error {
	service := rpc.NewLogUnitServiceClient(conn)
	logStream, err := service.TLogs(ctx, &rpc.TLogRQ{
		UnitId:         unitId,
		Duration:       duration,
		DateTimeLayout: dateTimeLayout,
		LogPattern:     logPattern,
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
			logging.StdErrLogger.Fatal(err)
		}

		result = append(result, rs.Log)
	}
	for i := len(result); i > 0; i = i - 1 {
		fmt.Println(result[i-1])
	}
	return nil
}
