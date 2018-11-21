#!/bin/bash
protoc -I . --proto_path=$GOPATH/src --go_out=plugins=grpc:. grpc.proto
