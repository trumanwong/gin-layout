#!/bin/bash

go install github.com/trumanwong/gin-project-gen@latest
gin-project-gen proto client api/app/v1/app.proto
gin-project-gen proto client api/app/v1/app_error.proto
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/modifier,sql/execquery,intercept ./internal/data/ent/schema
go install github.com/google/wire/cmd/wire@latest

cd cmd/server && wire && go build -o gin-layout-server && cd ../..

docker build -f Dockerfile -t docker.io/trumanwong/gin-layout-server:1.0.0 .