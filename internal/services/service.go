package services

import (
	"fmt"
	v1 "gin-layout/api/app/v1"
	"gin-layout/internal/conf"
	"gin-layout/internal/data"
	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
	"github.com/trumanwong/gin-transport/v2/transport/errors"
	"github.com/trumanwong/go-tools/cache"
	"github.com/trumanwong/go-tools/log"
	"github.com/trumanwong/go-tools/robot"
	"net/http"
)

type AppService struct {
	v1.UnimplementedAppServer
	greeterRepo *data.GreeterRepo

	//rabbitMQ       *mq.RabbitMQ
	cache         *cache.Cache
	config        *conf.Configs
	logger        *log.Logger
	snowflakeNode *snowflake.Node
	wechatRobot   *robot.WorkWechatRobot
}

func NewAppService(
	greeterRepo *data.GreeterRepo,

	cache *cache.Cache,
	config *conf.Configs,
	logger *log.Logger,
	wechatRobot *robot.WorkWechatRobot,
) *AppService {
	node, err := snowflake.NewNode(1)
	if err != nil {
		logger.Fatal("snowflake init error:", err)
		wechatRobot.SendText(&robot.SentTextRequest{
			Level:   robot.LevelError,
			Content: fmt.Sprintf("snowflake init error: %s", err.Error()),
		})
	}

	//rabbitMQ := mq.NewRabbitMQ(&mq.Options{
	//	Name: "",
	//	Addr: config.RabbitMQ.Url,
	//})
	return &AppService{
		greeterRepo: greeterRepo,

		cache:         cache,
		config:        config,
		logger:        logger,
		snowflakeNode: node,
		wechatRobot:   wechatRobot,
	}
}

func (s AppService) handleError(err error) {
	if err != nil {
		se := errors.FromError(err)
		if se.Code == http.StatusInternalServerError {
			s.wechatRobot.SendText(&robot.SentTextRequest{
				Level:   robot.LevelError,
				Content: err.Error(),
				IsAtAll: true,
			})
		}
		s.logger.Errorf(err.Error())
	}
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAppService)
