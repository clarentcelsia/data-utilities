package server

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9" // redis-client
)

type RedisInstance struct {
	Client *redis.Client
}

// Pre-requisite: Install your redis server
func (ins *RedisInstance) ConnectRedis(c context.Context) error {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // your redis server host
		Password: "",
		DB:       0, // redis db to partition data - separate services (caching, session, etc)
	})
	_, err := client.Ping(c).Result()
	if err != nil {
		return err
	}
	ins.Client = client
	return nil
}

// https://redis.io/commands/
func (ins *RedisInstance) SetDataToRedis(c context.Context, key string, val interface{}) error {
	cmd := ins.Client.Set(c, key, val, time.Duration(15)*time.Second)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (ins *RedisInstance) GetDataFromRedis(c context.Context, key string) (interface{}, error) {
	return ins.Client.Get(c, key).Result()
}
