package redisRepo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/viciousvs/blog-microservices/post/model/post"
	"sync"
)

var singleRedisDB *RedisDB
var initOnce sync.Once

func NewRedisDB(addr, password string, db int) *RedisDB {
	initOnce.Do(func() {
		ctx := context.Background()
		singleRedisDB = &RedisDB{}
		singleRedisDB.Conn = redis.NewClient(
			&redis.Options{
				Addr:     addr,
				Password: password,
				DB:       db,
			}).Conn(ctx)
	})
	return singleRedisDB
}

type RedisDB struct {
	Mu   sync.Mutex
	Conn *redis.Conn
}

func (r *RedisDB) Create(ctx context.Context, post post.Post) (string, error) {
	r.Mu.Lock()
	defer r.Mu.Unlock()
	err := r.Conn.HMSet(ctx, post.UUID, map[string]interface{
		post.
	}).Err()
}

func (r *RedisDB) GetAll(ctx context.Context) ([]*post.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisDB) GetById(ctx context.Context, uuid string) (post.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisDB) Update(ctx context.Context, post post.Post) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RedisDB) Delete(ctx context.Context, uuid string) (string, error) {
	//TODO implement me
	panic("implement me")
}
