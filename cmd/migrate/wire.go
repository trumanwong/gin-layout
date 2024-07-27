//go:build wireinject
// +build wireinject

package main

import (
	"gin-layout/internal/conf"
	"gin-layout/internal/data"
	"gin-layout/internal/robot"
	"github.com/google/wire"
)

func wireApp() *AppService {
	panic(wire.Build(
		conf.ProviderSet,
		robot.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
