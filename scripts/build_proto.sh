#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Remove all files from proto folder
rm -rf ../internal/proto_resources/*

# Generate protobuf files
protoc -I=../proto --go_out=../internal ../proto/*