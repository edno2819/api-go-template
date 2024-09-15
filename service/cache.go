package service

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

func GetDataCache(rdb *redis.Client, id string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cachedData, err := rdb.Get(ctx, id).Result()
	if err == redis.Nil {
		return ""
	}
	return cachedData
}
