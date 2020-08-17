VERSION=$(shell git describe --tags)
LDFLAGS=-ldflags "-s -w"

all: linux

release: all zip

clean:
	rm -rf bin/* *.zip

upx:
	upx -9 bin/*

linux:
	CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -o bin/server-linux-amd64 ${LDFLAGS} cmd/main.go
