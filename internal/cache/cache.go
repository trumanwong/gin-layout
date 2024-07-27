package cache

import (
	"gin-layout/internal/conf"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	cache2 "github.com/trumanwong/go-tools/cache"
	"log"
)

func NewCache(config *conf.Configs) *cache2.Cache {
	c, err := cache2.NewCache(
		&redis.Options{
			Addr:     config.Data.Redis.Addr,
			Password: config.Data.Redis.Password,
			DB:       int(config.Data.Redis.Db),
		},
		config.Data.Redis.Prefix,
	)
	if err != nil {
		log.Fatal("init cache error: ", err)
	}
	return c
}

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewCache)
