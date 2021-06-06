#!/bin/sh

protoc -I=proto --go_out=internal proto/msg.proto