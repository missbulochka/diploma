#!/usr/bin/env bash

# Generate Client and Server code for transactions-service using proto file
mkdir -p api/proto/gen
protoc -I=. \
    --go_out=. \
    --go-grpc_out=. \
    api/proto/transactions-service.proto