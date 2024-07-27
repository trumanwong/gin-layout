package robot

import (
	"gin-layout/internal/conf"
	"github.com/google/wire"
	"github.com/trumanwong/go-tools/robot"
)

func NewRobotWechat(config *conf.Configs) *robot.WorkWechatRobot {
	return robot.NewWorkWechatRobot(config.Other.Robot)
}

var ProviderSet = wire.NewSet(NewRobotWechat)
