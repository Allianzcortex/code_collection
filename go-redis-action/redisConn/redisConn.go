package redisconn

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func ConntectRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       15,
	})

	if _, err := conn.Ping(ctx).Result(); err != nil {
		log.Fatalf("Connect to redis failure because of %v\n", err)
	}

	return conn
}
