package deleteById

import (
	"context"
	"fmt"
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

func (h handler) Handle(ctx context.Context, request *pbPost.DeleteByIdRequest) (string, error) {
	uuid := request.GetUUID()
	if !utils.IsValidUUID(uuid) {
		return "", status.Errorf(codes.InvalidArgument, "Invalid UUID")
	}
	uuid, err := h.repo.Delete(ctx, uuid)
	if err != nil {
		return "", fmt.Errorf("err: %v", err)
	}
	return uuid, nil
}
