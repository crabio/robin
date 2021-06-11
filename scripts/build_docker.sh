#!/bin/sh

# Go to script directory
cd $(dirname $0)

# Go to root dir
cd ..

# Build docker image
docker build -t robin:latest .