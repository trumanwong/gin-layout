#!/bin/bash

protoc --proto_path=api/helloworld/v1  --proto_path=third_party --go_out=. --validate_out="lang=go:." api/helloworld/v1/greeter.proto
protoc --proto_path=api/error/v1 --go_out=. --validate_out="lang=go:." api/error/v1/error.proto
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier,sql/execquery,intercept ./internal/data/ent/schema
go install github.com/google/wire/cmd/wire@latest

cd cmd/server && wire && go build -o gin-layout-server && cd ../..

docker build -f Dockerfile -t docker.io/trumanwong/gin-layout-server:1.0.0 .