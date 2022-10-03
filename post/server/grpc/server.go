package grpc

import "github.com/viciousvs/blog-microservices/post/model/post"

type Server struct {
	repo post.Repository
}
