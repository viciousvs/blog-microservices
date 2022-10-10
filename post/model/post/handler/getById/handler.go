package getById

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post"
	"github.com/viciousvs/blog-microservices/post/utils"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type handler struct {
	repo post.Repository
}

func NewHandler(repo post.Repository) *handler {
	return &handler{repo: repo}
}

func (h handler) Handle(ctx context.Context, request *pbPost.GetByIdRequest) (post.Post, error) {
	uuid := request.GetUUID()
	if !utils.IsValidUUID(uuid) {
		return post.Post{}, status.Errorf(codes.InvalidArgument, "Invalid UUID")
	}
	return h.repo.GetById(ctx, uuid)
}
