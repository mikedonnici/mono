#!/bin/bash

PROJECT_ROOT="$1"
if [ -z "$1" ]
  then
    echo "No project root supplied"; exit 1
fi

echo "Generating Go proto code - project root: $PROJECT_ROOT"


# Note: module=github.com/mikedonnici/mono/go/... for go_opt and go-grpc_opt are required for
# outputting the generated code go to a location other than with the .proto files.
# Without this, the code is generated will the full path prefix github.com/mikedonnici/mono/go/...

OUT_BASE="$PROJECT_ROOT/go/internal"
GO_MODULE_PRE="github.com/mikedonnici/mono/go"
PROTO_BASE="$PROJECT_ROOT/proto"

# Each of these names should exist in the PROTO_BASE folder as <pkg>.proto
declare -a PackageNames=("attribute" "status")

for pkg in "${PackageNames[@]}"; do
  echo "Generating pkg '$pkg'"
  OUT="$OUT_BASE/$pkg"
  OPT="module=$GO_MODULE_PRE/$pkg"
  PROTO="$pkg.proto"
  protoc --go_out="$OUT" --go-grpc_out="$OUT" --go_opt="$OPT" --go-grpc_opt="$OPT" --proto_path="$PROTO_BASE" "$PROTO" --experimental_allow_proto3_optional
done
