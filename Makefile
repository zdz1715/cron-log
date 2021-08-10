GOOS="linux"
GOARCH="amd64"
BIN="./bin"

default: all

all: build

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BIN)/ .