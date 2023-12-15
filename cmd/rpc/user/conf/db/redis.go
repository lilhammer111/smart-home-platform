package db

import (
	"context"
	"git.zqbjj.top/pet/services/cmd/rpc/user/conf/binding"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	redisCli *redis.Client
	once     sync.Once
)

func GetRedis() *redis.Client {
	once.Do(initRedis)
	return redisCli
}

func initRedis() {
	redisCli = redis.NewClient(&redis.Options{
		Addr:     binding.GetRemoteConf().Redis.Addr,
		Username: binding.GetRemoteConf().Redis.Username,
		Password: binding.GetRemoteConf().Redis.Password,
		DB:       0,
	})
	if err = redisCli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
