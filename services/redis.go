package services

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func InitRedis() {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	Redis = redis.NewClient(opt)

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Panic(err)
	}

	log.Println("Redis connected")
}
