#!/bin/bash

PROJECT_ROOT="$1"
if [ -z "$1" ]
  then
    echo "No project root supplied"; exit 1
fi

echo "Generating Python proto code - project root: $PROJECT_ROOT"

OUT_BASE="$PROJECT_ROOT/py/pkg"
PROTO_BASE="$PROJECT_ROOT/proto"

# Each of these names should exist in the PROTO_BASE folder as <pkg>.proto
declare -a PackageNames=("attribute" "status")

for pkg in "${PackageNames[@]}"; do
  echo "Generating pkg '$pkg'"
  # This is just following poetry package/module naming convention - pkg-name/pkg_name
  OUT="$OUT_BASE"/"$pkg"-grpc/"$pkg"_grpc
  PROTO="$pkg.proto"
  python -m grpc_tools.protoc --python_out="$OUT" --grpc_python_out="$OUT" --proto_path="$PROTO_BASE" "$PROTO"
done
