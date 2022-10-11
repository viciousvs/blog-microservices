package add

import (
	"context"
	"github.com/viciousvs/blog-microservices/post/model/post"
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

func (h Handler) Handle(ctx context.Context, req *pbPost.CreateRequest) (string, error) {
	// validation 400
	if req.Title == "" {
		return "", status.Errorf(codes.InvalidArgument, "Empty Title")
	}
	// logic check 409
	//has post with uuid from request
	//h.repo.HasById // bool

	p := post.Post{
		UUID:    req.GetUUID(),
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}
	uuid, err := h.repo.Create(ctx, p)
	if err != nil {
		return "", status.Errorf(codes.AlreadyExists, "cannot create a post, err: %v", err)
	}
	// process run, create user throw repo

	return uuid, nil
}
