@echo off

SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -ldflags "-w -s" -o html_static_grpc main.go
upx html_static_grpc