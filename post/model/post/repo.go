package post

import "context"

type Repository interface {
	Create(ctx context.Context, post Post) (string, error)
	GetAll(ctx context.Context) ([]*Post, error)
	GetById(ctx context.Context, uuid string) (Post, error)
	Update(ctx context.Context, post Post) (string, error)
	Delete(ctx context.Context, uuid string) (string, error)
}
