package getById

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post"
	"github.com/viciousvs/blog-microservices/post/utils"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
)

type Handler struct {
	repo post.Repository
}

func NewHandler(repo post.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h Handler) Handle(ctx context.Context, request *pbPost.GetByIdRequest) (post.Post, error) {
	uuid := request.GetUUID()
	if !utils.IsValidUUID(uuid) {
		return post.Post{}, utils.ErrInvalidUUID
	}
	return h.repo.GetById(ctx, uuid)
}
