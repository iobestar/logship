#!/bin/bash
protoc -I ./unit/ --proto_path=./vendor --go_out=plugins=grpc:./unit/rpc ./unit/unit.proto
