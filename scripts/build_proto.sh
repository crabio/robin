#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Go to root project dir
cd ..
ls proto
# Generate protobuf files
protoc -I=proto --go_out=internal proto/*