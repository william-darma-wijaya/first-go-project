package config

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func NewRedisClient(viper *viper.Viper) *redis.Client {
    url := viper.GetString("REDIS_URL")
    if url == "" {
        panic("REDIS_URL is not set")
    }

    opt, err := redis.ParseURL(url)
    if err != nil {
        panic(err)
    }

    client := redis.NewClient(opt)

    ctx := context.Background()
    if err := client.Ping(ctx).Err(); err != nil {
        panic("failed to connect to redis: " + err.Error())
    } else {
		fmt.Println("REDIS CONNECTED!")
	}

    return client
}
