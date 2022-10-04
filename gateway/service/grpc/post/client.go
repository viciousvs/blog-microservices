package post

import (
	pbPost "github.com/viciousvs/blog-microservices/proto/post"
	"google.golang.org/grpc"
)

func Connect(address string) (*grpc.ClientConn, error) {
	return grpc.Dial(address, grpc.WithInsecure())
}

func GetClient(conn *grpc.ClientConn) pbPost.PostServiceClient {
	return pbPost.NewPostServiceClient(conn)
}
