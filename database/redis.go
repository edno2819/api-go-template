package database

import (
	"fmt"
	"gin-mongo-api/setting"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	url := setting.RedisSetting.Url
	fmt.Print(url)
	opts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}
