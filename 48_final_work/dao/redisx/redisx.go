package redisx

import (
	"fmt"
	"go_web_demo/48_final_work/settings"

	"github.com/go-redis/redis"
)

var (
	rdb *redis.Client
)

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%d", settings.ViperConfig.Redis.Host, settings.ViperConfig.Redis.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     dsn,
		Password: "",                                  // 密码
		DB:       settings.ViperConfig.Redis.Db,       // 数据库
		PoolSize: settings.ViperConfig.Redis.PoolSize, // 连接池大小
	})
	_, err = rdb.Ping().Result()
	return
}
