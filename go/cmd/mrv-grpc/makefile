.DEFAULT_GOAL := build
BINARY_NAME=mrv-grpc

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	# flags allow revive to pick up uncommented exported funcs
	golangci-lint run --exclude-use-default=false --enable=revive
.PHONY:lint

vet: fmt
	go vet ./...
.PHONY:vet

test: vet
	go test -v ./...

build: test
	go build -o ${BINARY_NAME} ./...
.PHONY:build

run: build
	./${BINARY_NAME}

clean:
	go clean

