#!/bin/bash

# Attribute
protoc --go_out=./internal/attribute \
	--go_opt=paths=source_relative \
	--go-grpc_out=./internal/attribute \
	--go-grpc_opt=paths=source_relative \
	--proto_path=../proto \
	attribute.proto

# Status
protoc --go_out=./internal/statuspb \
	--go_opt=paths=source_relative \
	--go-grpc_out=./internal/statuspb \
	--go-grpc_opt=paths=source_relative \
	--proto_path=../proto \
	status.proto
