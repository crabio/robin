#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Go to project root
cd ..

# Test
echo "Run unit tests"
go test -coverprofile=coverage.out -v ./... | exit 1

echo "Code coverage"
go tool cover -func coverage.out

# Lint
docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.40-alpine golangci-lint run
