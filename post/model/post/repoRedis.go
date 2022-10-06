package post

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/storage/redisRepo"
)

type repoRedis struct {
	redisSource redisRepo.RedisDB
}

func NewRepoRedis() Repository {
	return &repoRedis{redisSource: redisRepo.NewRedisDB()}
}

func (r repoRedis) Create(ctx context.Context, post Post) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r repoRedis) GetAll(ctx context.Context) ([]*Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r repoRedis) GetById(ctx context.Context, uuid string) (Post, error) {
	//TODO implement me
	panic("implement me")
}

func (r repoRedis) Update(ctx context.Context, post Post) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r repoRedis) Delete(ctx context.Context, uuid string) (string, error) {
	//TODO implement me
	panic("implement me")
}
