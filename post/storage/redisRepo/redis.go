package redisRepo

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/viciousvs/blog-microservices/post/config"
	"log"
	"sync"
	"time"
)

var singleRedisDB *RedisDB
var initOnce sync.Once

type RedisDB struct {
	*redis.Conn
}

func NewRedisDB(config config.RedisConfig) *RedisDB {
	initOnce.Do(func() {
		var err error
		singleRedisDB, err = newRedisRepository(config)
		if err != nil {
			log.Fatalf("cannot connect to Redis: %v", err)
		}
	})
	return singleRedisDB
}

func newRedisRepository(config config.RedisConfig) (*RedisDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn := redis.NewClient(
		&redis.Options{
			Addr:     config.Addr,
			Password: config.Password,
			DB:       config.DB,
		}).Conn(ctx)
	if err := conn.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("cannot ping redis: %v", err)
	}
	return &RedisDB{conn}, nil
}
