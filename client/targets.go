package client

import (
	"github.com/iobestar/logship/utils/logger"
	"google.golang.org/grpc"
)

type Targets map[string]*grpc.ClientConn

func CreateTargets(targets []string) Targets {

	if len(targets) == 0 {
		logger.Error.Fatalf("no targets")
	}

	result := make(Targets)
	for _, target := range targets {
		conn, err := grpc.Dial(target, grpc.WithInsecure())
		if err != nil {
			logger.Error.Fatalf("unable to create channel for target %s: %s", target, err.Error())
		}
		result[target] = conn
	}
	return result
}

func (tgc Targets) Close() {
	if len(tgc) > 0 {
		for _, channel := range tgc {
			if err := channel.Close(); nil != err {
				logger.Warning.Println(err)
			}
		}
	}
}