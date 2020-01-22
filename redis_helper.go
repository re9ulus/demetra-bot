package main

import (
	"github.com/go-redis/redis"
)

func GetNewRedisClient(addr string, passwd string) *redis.Client {
	return redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: passwd,
			DB:       0,
		})
}
