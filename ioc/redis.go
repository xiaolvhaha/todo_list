package ioc

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"strconv"
)

func InitRedis() redis.Cmdable {
	url := viper.GetString("redis.url")
	port := viper.GetInt("redis.port")

	return redis.NewClient(&redis.Options{
		Addr: url + ":" + strconv.Itoa(port),
	})
}
