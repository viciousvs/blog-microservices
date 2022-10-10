package post

import (
	"fmt"
	"github.com/viciousvs/blog-microservices/gateway/config"
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc"
	"log"
)

type client struct {
}

func NewPostClientGRPC(cfg config.PostClientConfig) (pbPost.PostServiceClient, error) {
	conn, err := grpc.Dial(cfg.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to grpc Server, err:%v", err)
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	return pbPost.NewPostServiceClient(conn), nil
}
