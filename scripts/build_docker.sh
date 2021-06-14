#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Go to root dir
cd ..

# Build docker image for Linux
docker build -t robin:latest .
