package main

import (
	"fmt"
	docs "gin-layout/cmd/server/docs"
	"gin-layout/internal/conf"
	"gin-layout/internal/middlewares"
	"gin-layout/internal/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/trumanwong/go-tools/log"
	toolMiddlewares "github.com/trumanwong/go-tools/middlewares"
)

type AppService struct {
	router *gin.Engine
	port   uint64
}

func newApp(
	s *services.AppService,
	config *conf.Configs,
	logger *log.Logger,
	middleware *middlewares.Middleware,
) *AppService {
	trace := toolMiddlewares.NewTracing(nil, logger)

	gin.SetMode(config.Server.Http.Mode)
	allowCors := toolMiddlewares.NewAllowCors(
		config.Server.Http.Mode,
		"Content-Type,Accept,Authorization,X-Requested-With,X-XSRF-TOKEN,x-csrf-token,Cache-Control",
		"POST, GET, OPTIONS, PUT, DELETE, PATCH", config.Other.AllowOrigins,
	)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(trace.Handle())
	router.Use(allowCors.Handle())
	// swagger文档
	if config.Server.Http.Mode != gin.ReleaseMode {
		docs.SwaggerInfo.BasePath = "/api/v1"
		// 非生产环境才开启swagger
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
	// 非必要鉴权的api
	unAuthApiV1(router, middleware, s)
	// 需要鉴权的api
	authApiV1(router, middleware, s)
	//staticApi(router, config)

	return &AppService{
		router: router,
		port:   config.Server.Http.Port,
	}
}

// 不需要鉴权
func unAuthApiV1(router *gin.Engine, middleware *middlewares.Middleware, s *services.AppService) {
	unAuthApi := router.Group("/api/v1")
	unAuthApi.Use(middleware.UnAuthCheck())

	greeterApi := unAuthApi.Group("greeter")
	{
		// 获取标签列表
		greeterApi.POST("", s.CreateGreeter)
	}
}

// 需要鉴权
func authApiV1(router *gin.Engine, middleware *middlewares.Middleware, s *services.AppService) {
	authApi := router.Group("/api/v1")
	authApi.Use(middleware.AuthCheck())
}

// 静态路由
//func staticApi(router *gin.Engine, config *conf.Configs) {
//	router.Static("/storage", config.Other.StoragePath)
//}

func main() {
	app := wireApp()
	err := app.router.Run(fmt.Sprintf(":%d", app.port))
	if err != nil {
		panic(err)
	}
}
