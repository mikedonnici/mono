.DEFAULT_GOAL := test

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
	docker-compose -f docker-compose.test.yml up --detach
	go test -v ./...
	docker-compose -f docker-compose.test.yml down --volume

clean:
	go clean

