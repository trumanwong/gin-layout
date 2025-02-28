package main

import (
	"fmt"
	v1 "gin-layout/api/app/v1"
	"gin-layout/internal/conf"
	"gin-layout/internal/middlewares"
	"gin-layout/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/trumanwong/gin-transport/v2/transport"
	"github.com/trumanwong/go-tools/log"
	toolMiddlewares "github.com/trumanwong/go-tools/middlewares"
)

func NewHttpServer(
	config *conf.Configs,
	logger *log.Logger,
	middleware *middlewares.Middleware,
	srv *services.AppService,
) *transport.Server {
	gin.SetMode(config.Server.Http.Mode)
	trace := toolMiddlewares.NewTracing(nil, logger)

	allowCors := toolMiddlewares.NewAllowCors(
		config.Server.Http.Mode,
		"Content-Type,Accept,Authorization,X-Requested-With,X-XSRF-TOKEN,x-csrf-token,Cache-Control",
		"POST, GET, OPTIONS, PUT, DELETE, PATCH", config.Other.AllowOrigins,
	)
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(trace.Handle())
	engine.Use(allowCors.Handle())
	server := transport.NewServer(engine, []string{
		fmt.Sprintf(":%d", config.Server.Http.Port),
	}, []transport.Middleware{
		middleware.AuthCheck,
	})

	v1.RegisterAppHTTPServer(server, srv)
	return server
}

func main() {
	server := wireApp()
	err := server.Run()
	if err != nil {
		panic(err)
	}
}
