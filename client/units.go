package client

import (
	"context"
	"fmt"
	"github.com/iobestar/logship/unit/rpc"
	"io"
)

func Units(ctx context.Context, targets Targets) error {
	for _, target := range targets {
		service := rpc.NewLogUnitServiceClient(target)
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
	}
	return nil
}