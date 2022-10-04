package main

import (
	"github.com/viciousvs/blog-microservices/post/model/post"
	postGRPC "github.com/viciousvs/blog-microservices/post/server/grpc"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	// init env
	// init server
	inMemRepo := post.NewInMemRepo()
	inMemRepo.DB = append(inMemRepo.DB, &post.Post{UUID: "awlfkmawfmakwfawf", Title: "wafawfawf", Content: "wfawfawmfwaofm", CreatedAt: time.Now(), UpdateAt: time.Now()})
	srv := postGRPC.NewServer(inMemRepo)

	listener, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Cannot listen: %v", err)
	}
	defer func() {
		err := listener.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	s := grpc.NewServer()
	pbPost.RegisterPostServiceServer(s, srv)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Cannot serve a grpc server: %v", err)
	}
}
