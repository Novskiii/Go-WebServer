.PHONY: all build
BINARY="awosone"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./main.go conf/config.yaml

gotool:
	go fmt ./
	go vet ./
