// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"gin-layout/internal/conf"
	"gin-layout/internal/data"
	"gin-layout/internal/robot"
)

// Injectors from wire.go:

func wireApp() *AppService {
	configs := conf.NewConfig()
	client := data.NewEntClient(configs)
	dataData := data.NewData(client)
	workWechatRobot := robot.NewRobotWechat(configs)
	appService := newApp(dataData, workWechatRobot)
	return appService
}
