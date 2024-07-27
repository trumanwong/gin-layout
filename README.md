# Gin Project Template

## Install gin-layout-gen
```
go install github.com/trumanwong/gin-layout-gen@latest
```
## Create a gin project
```
# Create a template project
gin-layout-gen new project_name
cd project_name
# generate swagger
/bin/bash ./swagger.sh
# start server
go run cmd/server/main.go cmd/server/wire_gen.go
```

## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```