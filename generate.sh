#!/bin/bash

PROTO_DIR="proto"
OUT_DIR="./gen/go"
find "$PROTO_DIR" -name "*.proto" | while read -r proto_file; do
    echo "Processing $proto_file..."
    protoc -I "$PROTO_DIR" "$proto_file" --go_out="$OUT_DIR" --go_opt=paths=source_relative --go-grpc_out="$OUT_DIR" --go-grpc_opt=paths=source_relative
    if [ $? -ne 0 ]; then
        echo "Failed to process $proto_file"
        exit 1
    fi
done

echo "All .proto files have been processed."
