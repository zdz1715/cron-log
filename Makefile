GOOS="linux"
GOARCH="amd64"
BIN="./bin"

default: all

all: build

build:
	GO111MODULE=on GOPROXY=https://mirrors.aliyun.com/goproxy/,direct GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN)/ .