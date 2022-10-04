package grpc

import (
	"github.com/viciousvs/blog-microservices/post/model/post"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
)

type Server struct {
	pbPost.UnimplementedPostServiceServer
	repo post.Repository
}

func NewServer(repo post.Repository) *Server {
	return &Server{repo: repo}
}
