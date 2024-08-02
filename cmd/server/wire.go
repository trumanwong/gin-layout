//go:build wireinject
// +build wireinject

package main

import (
	"gin-layout/internal/cache"
	"gin-layout/internal/conf"
	"gin-layout/internal/data"
	"gin-layout/internal/log"
	"gin-layout/internal/middlewares"
	"gin-layout/internal/robot"
	"gin-layout/internal/services"
	"github.com/google/wire"
	"github.com/trumanwong/gin-transport/transport"
)

func wireApp() *transport.Server {
	panic(wire.Build(
		cache.ProviderSet,
		conf.ProviderSet,
		log.ProviderSet,
		robot.ProviderSet,
		data.ProviderSet,
		middlewares.ProviderSet,
		services.ProviderSet,
		NewHttpServer,
	))
}
