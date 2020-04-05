#!/bin/bash
protoc -I ./unit/ --go_out=plugins=grpc:./unit/rpc ./unit/unit.proto
