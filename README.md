# Gin Project Template

## Install gin-project-gen
```
go install github.com/trumanwong/gin-project-gen@latest
```
## Create a gin project
```
# Create a template project
gin-project-gen new project_name
cd project_name
# start server
go run cmd/server/main.go cmd/server/wire_gen.go -config deploy/configs/config.yaml
```

## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## License

This project is licensed under the MIT License - see the [LICENSE.md](https://github.com/trumanwong/gin-layout/blob/main/LICENSE) file for details.

## Acknowledgements

The following project had particular influence on protoc-gen-go-gin's design.

- [go-kratos/kratos](https://github.com/go-kratos/kratos) is a microservice-oriented governance framework implemented by golang.
- [gin-gonic/gin](https://github.com/gin-gonic/gin) is a web framework written in Go.
- <a href="https://jb.gg/OpenSourceSupport"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.svg?_gl=1*1nuywz*_ga*NTcwMDkwNDIxLjE2ODQzMTI1Mzg.*_ga_9J976DJZ68*MTY4NDMxMjUzOC4xLjEuMTY4NDMxMjU1Mi4wLjAuMA.." width="60" height="60"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand.svg?_gl=1*1nuywz*_ga*NTcwMDkwNDIxLjE2ODQzMTI1Mzg.*_ga_9J976DJZ68*MTY4NDMxMjUzOC4xLjEuMTY4NDMxMjU1Mi4wLjAuMA.." width="60" height="60"><img src="https://resources.jetbrains.com/storage/products/company/brand/logos/GoLand_icon.svg?_gl=1*1b2zdbh*_ga*NTcwMDkwNDIxLjE2ODQzMTI1Mzg.*_ga_9J976DJZ68*MTY4NDMxMjUzOC4xLjEuMTY4NDMxMjU1Mi4wLjAuMA.." width="60" height="60"></a>
