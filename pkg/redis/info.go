package redis

import (
	"context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-redis/redis/v8"
	"redis_mate_go/pkg/config"
)

func Info() string {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Hostname,
		Password: config.Password, // no password set
	})
	defer rdb.Close()
	info := rdb.Info(context.Background())
	if info.Err() != nil {
		logs.Error("get redis info error ", info.Err())
		return ""
	}
	return info.Val()
}
