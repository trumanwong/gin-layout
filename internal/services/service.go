package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	errorV1 "gin-layout/api/error/v1"
	"gin-layout/internal/conf"
	"gin-layout/internal/data"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/trumanwong/cryptogo"
	"github.com/trumanwong/cryptogo/paddings"
	"github.com/trumanwong/go-tools/cache"
	"github.com/trumanwong/go-tools/log"
	"github.com/trumanwong/go-tools/robot"
	"net/http"
)

type AppService struct {
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

func (s AppService) handleError(ctx *gin.Context, err error, statusCode int) {
	if err != nil {
		if statusCode == http.StatusInternalServerError {
			s.wechatRobot.SendText(&robot.SentTextRequest{
				Level:   robot.LevelError,
				Content: err.Error(),
				IsAtAll: true,
			})
		}
		s.logger.Errorf(err.Error())
		s.Response(ctx, statusCode, errorV1.Error{
			Message: err.Error(),
		})
	}
}

func (s AppService) Response(ctx *gin.Context, code int, data interface{}) {
	if !s.config.Crypto.Enable || ctx.GetHeader(s.config.Crypto.PlainHeaderKey) == s.config.Crypto.PlainHeaderVal {
		ctx.JSON(code, gin.H{
			"data": data,
		})
		return
	}
	plaintext, _ := json.Marshal(data)
	cipher, _ := cryptogo.AesCBCEncrypt(plaintext, []byte(s.config.Crypto.AesKey), []byte(s.config.Crypto.AesIv), paddings.PKCS7)
	ctx.JSON(code, gin.H{
		"data": base64.StdEncoding.EncodeToString(cipher),
	})
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAppService)
