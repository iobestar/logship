package client

import (
	"context"
	"fmt"
	"github.com/iobestar/logship/unit/rpc"
	"google.golang.org/grpc"
	"io"
)

func Units(ctx context.Context, conn *grpc.ClientConn) error {
	service := rpc.NewLogUnitServiceClient(conn)
	unitStream, err := service.GetUnits(ctx, &rpc.Empty{})
	if nil != err {
		return err
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
	return nil
}
