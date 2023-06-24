#!/bin/bash

mkdir ./proto_services

protoc --go_out=./proto/services   \
    --go-grpc_out=./proto/services \
    proto/doggo-proto/*