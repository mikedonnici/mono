.DEFAULT_GOAL := build
BINARY_NAME=grpc-srvr

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

build:
	go build -o ${BINARY_NAME} ./...
.PHONY:build

run: build
	./${BINARY_NAME} $(ARGS)

clean:
	go clean

# shortcuts
run-sample: build
	./${BINARY_NAME} -cfg testdata/sample.cfg

test-run-sample: test run-sample

