#!/bin/bash

VERSION=$1
GIT_COMMIT=$(git rev-list -1 HEAD)
rm -f logship
GOOS=$os GOARCH=amd64 go build -ldflags "-X main.Version=$VERSION -X main.Revision=$GIT_COMMIT -s -w"  ./cmd/logship
upx --brute ./logship
docker rmi logship:$VERSION
docker build . -t logship:$VERSION
rm -f logship
