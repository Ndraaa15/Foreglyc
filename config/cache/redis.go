package cache

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func New() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("cache.address"), viper.GetInt("cache.port")),
		Password: viper.GetString("cache.password"),
		DB:       viper.GetInt("cache.db"),
		Username: viper.GetString("cache.username"),
	})

	return rdb
}
