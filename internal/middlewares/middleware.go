package middlewares

import (
	"gin-layout/internal/conf"
	"github.com/google/wire"
	"github.com/trumanwong/go-tools/log"
	"github.com/trumanwong/go-tools/robot"
)

type Middleware struct {
	config      *conf.Configs
	logger      *log.Logger
	wechatRobot *robot.WorkWechatRobot
}

func NewMiddleware(
	config *conf.Configs,
	logger *log.Logger,
	wechatRobot *robot.WorkWechatRobot,
) *Middleware {
	return &Middleware{
		config:      config,
		logger:      logger,
		wechatRobot: wechatRobot,
	}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewMiddleware)
