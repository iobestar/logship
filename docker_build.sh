#!/bin/bash

VERSION=$1

rm -f logship
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" ./cmd/logship
docker rmi logship:$VERSION
docker build . -t logship:$VERSION
rm -f logship
