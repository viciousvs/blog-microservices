package update

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post"
)

type handler struct {
	repo post.Repository
}

func NewHandler(repo post.Repository) *handler {
	return &handler{repo: repo}
}

func (h handler) Handle(ctx context.Context, post post.Post) (string, error) {
	return h.repo.Update(ctx, post)
}
