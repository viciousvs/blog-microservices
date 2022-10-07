package post

import (
	"context"
	UUID "github.com/google/uuid"
	"github.com/viciousvs/blog-microservices/post/utils"
	"sync"
	"time"
)

type MemDB struct {
	Mu sync.Mutex
	DB []*Post
}

func (m *MemDB) GetNewUUID() string {
	uuid := UUID.New()
	return uuid.String()
}

func NewInMemRepo() *MemDB {
	return &MemDB{Mu: sync.Mutex{}, DB: make([]*Post, 0)}
}
func (m *MemDB) Create(ctx context.Context, post Post) (string, error) {
	if m.DB == nil {
		return "", utils.ErrNilDB
	}
	m.Mu.Lock()
	defer m.Mu.Unlock()

	now := time.Now().Unix()
	post = Post{
		UUID:      m.GetNewUUID(),
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}

	m.DB = append(m.DB, &post)
	return post.UUID, nil
}

func (m *MemDB) GetAll(ctx context.Context) ([]*Post, error) {
	if m.DB == nil {
		return make([]*Post, 0), utils.ErrNilDB
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()
	return m.DB, nil
}

func (m *MemDB) GetById(ctx context.Context, uuid string) (Post, error) {
	if m.DB == nil {
		return Post{}, utils.ErrNilDB
	}
	if len(uuid) == 0 {
		return Post{}, utils.ErrEmptyUUID
	}
	m.Mu.Lock()
	defer m.Mu.Unlock()
	var post Post
	var isExist bool
	for _, p := range m.DB {
		if p.UUID == uuid {
			post = *p
			isExist = true
		}
	}

	if !isExist {
		return post, utils.ErrNotExist
	}
	return post, nil
}

func (m *MemDB) Update(ctx context.Context, post Post) (string, error) {
	if m.DB == nil {
		return "", utils.ErrNilDB
	}
	var uuid string
	var isExist bool
	for k, p := range m.DB {
		if p.UUID == post.UUID {
			m.DB[k] = &Post{
				UUID:      post.UUID,
				Title:     post.Title,
				Content:   post.Content,
				UpdatedAt: time.Now().Unix(),
			}
			uuid = p.UUID
			isExist = true
		}
	}
	if !isExist {
		return "", utils.ErrNotExist
	}
	return uuid, nil
}

func (m *MemDB) Delete(ctx context.Context, uuid string) (string, error) {
	if m.DB == nil {
		return "", utils.ErrNilDB
	}
	if len(uuid) == 0 {
		return "", utils.ErrEmptyUUID
	}
	var isExist bool
	var index int
	for indx, p := range m.DB {
		if p.UUID == uuid {
			isExist = true
			index = indx
		}
	}

	if !isExist {
		return "", utils.ErrNotExist
	}

	m.DB = append(m.DB[:index], m.DB[index+1:]...)
	return uuid, nil
}
