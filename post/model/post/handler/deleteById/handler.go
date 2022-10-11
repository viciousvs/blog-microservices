package deleteById

import (
	"context"
	"fmt"
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

func (h Handler) Handle(ctx context.Context, request *pbPost.DeleteByIdRequest) (string, error) {
	uuid := request.GetUUID()
	if !utils.IsValidUUID(uuid) {
		return "", utils.ErrInvalidUUID
	}
	uuid, err := h.repo.Delete(ctx, uuid)
	if err != nil {
		return "", fmt.Errorf("err: %v", err)
	}
	return uuid, nil
}
