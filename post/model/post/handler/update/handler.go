package update

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post"
	"github.com/viciousvs/blog-microservices/post/utils"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	repo post.Repository
}

func NewHandler(repo post.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h Handler) Handle(ctx context.Context, req *pbPost.UpdateRequest) (string, error) {
	if !utils.IsValidUUID(req.UUID) {
		return "", status.Errorf(codes.InvalidArgument, "Invalid UUID")
	}
	// logic check 409

	p := post.Post{
		UUID:    req.GetUUID(),
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}
	uuid, err := h.repo.Update(ctx, p)
	if err != nil {
		if err == utils.ErrNotFound {
			return "", err
		}
		return "", utils.ErrNotingUpdate
	}
	// process run, create user throw repo

	return uuid, nil
}
