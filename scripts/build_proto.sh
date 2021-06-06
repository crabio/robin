#!/bin/sh

# Go to script directory
cd $(dirname $0)

cd ..

# Generate protobuf files
protoc -I=proto --go_out=internal proto/msg.proto