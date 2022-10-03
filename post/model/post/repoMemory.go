package post

import (
	"github.com/viciousvs/blog-microservices/post/utils"
	"sync"
)

type MemDB struct {
	Mu *sync.Mutex
	DB []*Post
}

func (m *MemDB) Create(post Post) (string, error) {
	if m.Mu == nil {
		return "", utils.ErrNilMutex
	}
	if m.DB == nil {
		return "", utils.ErrNilDB
	}
	m.Mu.Lock()
	defer m.Mu.Unlock()

	m.DB = append(m.DB, &post)
	return post.UUID, nil
}

func (m MemDB) GetAll() ([]*Post, error) {
	if m.Mu == nil {
		return make([]*Post, 0), utils.ErrNilMutex
	}
	if m.DB == nil {
		return make([]*Post, 0), utils.ErrNilDB
	}

	m.Mu.Lock()
	defer m.Mu.Unlock()
	return m.DB, nil
}

func (m MemDB) GetById(uuid string) (Post, error) {
	if m.Mu == nil {
		return Post{}, utils.ErrNilMutex
	}
	if m.DB == nil {
		return Post{}, utils.ErrNilDB
	}
	if len(uuid) == 0 {
		return Post{}, utils.ErrEmptyUUID
	}
	m.Mu.Lock()
	defer m.Mu.Unlock()
	var post Post
	for _, p := range m.DB {
		if p.UUID == uuid {
			post = *p
		}
	}
	return post, nil
}

func (m MemDB) Update(post Post) (string, error) {
	if m.Mu == nil {
		return "", utils.ErrNilMutex
	}
	if m.DB == nil {
		return "", utils.ErrNilDB
	}
	var uuid string
	for k, p := range m.DB {
		if p.UUID == post.UUID {
			m.DB[k] = &post
			uuid = p.UUID
		}
	}
	return uuid, nil
}

func (m *MemDB) name()  {

}
func (m MemDB) Delete(uuid string) (string, error) {
	if m.Mu == nil {
		return "", utils.ErrNilMutex
	}
	if m.DB == nil {
		return "", utils.ErrNilDB
	}
	if len(uuid) == 0 {
		return "", utils.ErrEmptyUUID
	}
	var
	for _, p := range m.DB {
		if p.UUID == uuid {
			post = *p
		}
	}
	return post, nil
}
