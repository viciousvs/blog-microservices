package post

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	UUID "github.com/google/uuid"
	"github.com/viciousvs/blog-microservices/post/config"
	"github.com/viciousvs/blog-microservices/post/storage/redisRepo"
	"github.com/viciousvs/blog-microservices/post/utils"
	"log"
	"strconv"
	"strings"
	"time"
)

const postPrefix = "POST*"

func (r *repoRedis) getNewUUID() string {
	uuid := UUID.New()
	return uuid.String()
}

type repoRedis struct {
	redisSource *redisRepo.RedisDB
}

func NewRepoRedis(config config.RedisConfig) Repository {
	return &repoRedis{redisSource: redisRepo.NewRedisDB(config)}
}

func (r repoRedis) Create(ctx context.Context, post Post) (string, error) {
	now := time.Now().Unix()
	uuid := r.getNewUUID()
	err := r.redisSource.HMSet(ctx, post.UUID, map[string]interface{}{
		"uuid":      uuid,
		"title":     post.Title,
		"content":   post.Content,
		"createdAt": now,
		"updatedAt": now,
	}).Err()
	if err != nil {
		return "", fmt.Errorf("cannot set a post to redis: %v", err)
	}
	return uuid, nil
}

func (r repoRedis) GetAll(ctx context.Context) ([]*Post, error) {
	posts := make([]*Post, 0)

	postKeys, err := r.redisSource.Keys(ctx, postPrefix).Result()
	if err != nil {
		return posts, fmt.Errorf("Cannot get keys from redis by prefix:%s, err:%v", postPrefix, err)
	}

	for _, key := range postKeys {
		m, err := r.redisSource.HGetAll(ctx, key).Result()
		if err != nil {
			continue
		}
		p := parsePost(key, m)
		posts = append(posts, &p)
	}

	return posts, nil
}

func (r repoRedis) GetById(ctx context.Context, uuid string) (Post, error) {
	var p Post
	key := postPrefix + uuid
	m, err := r.redisSource.HGetAll(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return p, utils.ErrNotExist
		}
		return p, fmt.Errorf("cannot hash get all, err:%v", err)
	}
	p = parsePost(key, m)
	return p, nil
}

func (r repoRedis) Update(ctx context.Context, inputPost Post) (string, error) {
	//TODO fix Update Method
	var p Post
	//tmpCreatedAt := inputPost.CreatedAt
	key := postPrefix + inputPost.UUID
	m, err := r.redisSource.HGetAll(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", utils.ErrNotExist
		}
		return "", fmt.Errorf("cannot hash get all, err:%v", err)
	}

	p = parsePost(key, m)

	if strings.Trim(inputPost.Title, " ") != "" {
		p.Title = inputPost.Title
	}
	if strings.Trim(inputPost.Content, "") != "" {
		p.Content = inputPost.Content
	}
	//TODO UpdatedAt in if or out
	p.UpdatedAt = time.Now().Unix()

	err = r.redisSource.HMSet(ctx, key, map[string]interface{}{
		"uuid":      p.UUID,
		"title":     p.Title,
		"content":   p.Content,
		"createdAt": p.CreatedAt,
		"updatedAt": p.UpdatedAt,
	}).Err()
	if err != nil {
		return "", fmt.Errorf("cannot update a post in redis: %v", err)
	}
	return p.UUID, nil
}

func (r repoRedis) Delete(ctx context.Context, uuid string) (string, error) {
	key := postPrefix + uuid
	n, err := r.redisSource.Del(ctx, key).Result()
	if n != 1 {
		return "", utils.ErrNotExist
	}
	if err != nil {
		return "", fmt.Errorf("cannot delete from redis, err:%v", err)
	}
	return uuid, nil
}

func parsePost(key string, pMap map[string]string) Post {
	uuid := strings.TrimLeft(key, postPrefix)
	createdAt, err := strconv.ParseInt(pMap["createdAt"], 10, 64)
	if err != nil {
		log.Printf("cannot parse int, key:%s, err:%v", key, err)
	}
	updatedAt, err := strconv.ParseInt(pMap["updatedAt"], 10, 64)
	if err != nil {
		log.Printf("cannot parse int, key:%s, err:%v", key, err)
	}
	return Post{
		UUID:      uuid,
		Title:     pMap["title"],
		Content:   pMap["content"],
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}
