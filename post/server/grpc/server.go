package grpc

import (
	"fmt"
	"github.com/viciousvs/blog-microservices/post/config"
	"github.com/viciousvs/blog-microservices/post/model/post"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pbPost.UnimplementedPostServiceServer
	repo post.Repository
}

func NewServer(repo post.Repository) *Server {
	return &Server{repo: repo}
}

func (s *Server) Run(cfg config.ServerConfig) error {
	listener, err := net.Listen("tcp", cfg.Addres)
	if err != nil {
		return fmt.Errorf("Cannot listen: %v", err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	srv := grpc.NewServer()
	pbPost.RegisterPostServiceServer(srv, s)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("Cannot serve a grpc server: %v", err)
	}
	return nil
}
