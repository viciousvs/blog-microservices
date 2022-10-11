package getAll

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post"
)

type Handler struct {
	repo post.Repository
}

func NewHandler(repo post.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h Handler) Handle(ctx context.Context) ([]*post.Post, error) {
	return h.repo.GetAll(ctx)
}
