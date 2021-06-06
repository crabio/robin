#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Generate protobuf files
protoc -I=../proto --go_out=../internal ../proto/*