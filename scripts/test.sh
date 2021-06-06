#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Go to project root
cd ..

# Test
go test go test -v ./...

# Lint
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.40-alpine golangci-lint run