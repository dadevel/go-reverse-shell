#!/usr/bin/env bash
set -eu

export CGO_ENABLED=0 GOARCH=amd64 GOOS=linux "$@"
go build -o "./shell-$GOOS-$GOARCH" -ldflags "-extldflags -static -s -w -X main.Shell=${shell@Q} -X main.Destination=${destination@Q}" ./shell.go

