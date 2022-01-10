#!/bin/bash

PROJECT_ROOT="$1"
if [ -z "$1" ]
  then
    echo "No project root supplied"; exit 1
fi

echo "Generating protobuf things (root: $PROJECT_ROOT)"

#protoc --go_out="$PROJECT_ROOT"/go/internal/attribute \
#	--go_opt=paths=source_relative \
#	--go-grpc_out="$PROJECT_ROOT"/go/internal/attribute \
#	--go-grpc_opt=paths=source_relative \
#	--proto_path="$PROJECT_ROOT"/proto \
#	attribute.proto


# Attribute
protoc --go_out="$PROJECT_ROOT"/go/internal/attribute \
	--go-grpc_out="$PROJECT_ROOT"/go/internal/attribute \
	--proto_path="$PROJECT_ROOT"/proto \
	attribute.proto

# Status
protoc --go_out="$PROJECT_ROOT"/go/internal/status \
	--go-grpc_out="$PROJECT_ROOT"/go/internal/status \
	--proto_path="$PROJECT_ROOT"/proto \
	status.proto
