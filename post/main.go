package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/viciousvs/blog-microservices/post/config"
	"github.com/viciousvs/blog-microservices/post/model/post"
	postGRPC "github.com/viciousvs/blog-microservices/post/server/grpc"
	"log"
)

func init() {
	// loads values from post.env into the system
	if err := godotenv.Load("post.env"); err != nil {
		log.Print("No post.env file found")
	}
}
func main() {
	// init env
	cfg := config.NewConfig()
	// init server
	repo := post.NewRepoRedis(cfg.Redis)

	fmt.Println(cfg.Server.Addres)
	s := postGRPC.NewServer(repo)
	if err := s.Run(cfg.Server); err != nil {
		log.Fatalf("server not started, %v", err)
	}
}
