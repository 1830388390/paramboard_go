package model

import (
	"context"
	"github.com/go-redis/redis"
)

// MQ rabbitMQ链接单例
var _rd *redis.Client

// RabbitMQ 在中间件中初始化rabbitMQ链接
func Redis(addr, passwprd *string, DB *int) {
	conn := redis.NewClient(&redis.Options{
		Addr:     *addr,
		Password: *passwprd,
		DB:       *DB,
	})
	_, err := conn.Ping().Result()
	if err != nil {
		panic(err)
	}
	_rd = conn
}

func NewRedisClient(ctx context.Context) *redis.Client {
	rd := _rd
	return rd
}
